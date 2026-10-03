package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/arl/statsviz"
	"github.com/gorilla/mux"
	"github.com/rajat-mangla/go-boilerplate/config"
	"github.com/rs/zerolog/log"
	"github.com/tylerb/graceful"
)

// HTTPAPIServer is an HTTP server for serving APIs.
// This server comes with some default bells & whistles for
// - Logging
// - HTTP Context
// - HTTP Metrics
// - New Relic integration
// - Health Checks
// - Debug/Profile endpoints
// These features are configurable using middlewares.
type HTTPAPIServer struct {
	server         *graceful.Server
	statsvizServer *statsviz.Server
}

// NewHTTPAPIServer creates a new HTTPAPIServer using a mux.Router.
func NewHTTPAPIServer(cfg config.HTTPAPIConfig, r *mux.Router) *HTTPAPIServer {
	// always attach a health check endpoint
	r.Methods("GET").Path("/ping").HandlerFunc(internalPingHandler)

	// attach pprof & statsviz in debug mode
	var statsvizServer *statsviz.Server
	if cfg.DebugMode {
		statsvizServer, _ = statsviz.NewServer()

		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)
		r.HandleFunc("/debug/pprof/goroutine", pprof.Index)
		r.HandleFunc("/debug/pprof/heap", pprof.Index)
		r.HandleFunc("/debug/pprof/threadcreate", pprof.Index)
		r.Methods("GET").Path("/debug/statsviz/ws").Name("GET /debug/statsviz/ws").HandlerFunc(statsvizServer.Ws())
		r.Methods("GET").PathPrefix("/debug/statsviz/").Name("GET /debug/statsviz/").Handler(statsvizServer.Index())
	}

	// create the http server
	httpServer := http.Server{
		Addr:    cfg.ListenAddr,
		Handler: r,
	}

	// ... and wrap it for graceful shutdowns.
	// graceful shutdowns will ensure abrupt closing/stopping
	// of ongoing requests. Even though, a HAProxy will drain
	// connection. This adds another safety net to ensure critical
	// writes are not missed.
	server := &graceful.Server{
		Timeout: 15 * time.Second,
		Server:  &httpServer,
	}

	return &HTTPAPIServer{
		server:         server,
		statsvizServer: statsvizServer,
	}
}

// Start starts the API server in a new goroutine.
func (a *HTTPAPIServer) Start() {
	go func(a *HTTPAPIServer) {
		log.Info().Msgf("[API.SERVER] PID %d. Starting server on %s", os.Getpid(), a.server.Addr)

		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("[API.SERVER] Unhandled server shutdown: %s", err.Error())
		}
	}(a)
}

// Shutdown stops the HTTP server.
func (a *HTTPAPIServer) Shutdown() error {
	shutdownErr := a.server.Shutdown(context.Background())
	if a.statsvizServer == nil {
		return shutdownErr
	}
	return errors.Join(shutdownErr, a.statsvizServer.Close())
}

// WaitForShutdown blocks till the server is shutdown.
// Shutdowns can be initiated with a SIGHUP, SIGKILL interupts.
func (a *HTTPAPIServer) WaitForShutdown() bool {
	<-a.server.StopChan()
	log.Info().Msgf("[API.SERVER] Server shutdown complete")

	return true
}

func internalPingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(("OK")))
}
