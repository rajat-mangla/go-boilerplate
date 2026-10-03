package main

import (
	"github.com/gorilla/mux"
	"github.com/rajat-mangla/go-boilerplate/config"
	"github.com/rajat-mangla/go-boilerplate/handlers"
	"github.com/rajat-mangla/go-boilerplate/service"
)

type RouterOptionFunc func(*mux.Router)

func NewRouterWithOptions(options ...RouterOptionFunc) *mux.Router {
	r := mux.NewRouter()
	for _, option := range options {
		option(r)
	}
	return r
}

func WithInternalRoutes(cfg *config.Config, sr *service.Registry) RouterOptionFunc {
	return func(r *mux.Router) {
		addInternalRoutes(r, *cfg, sr)
	}
}

func addInternalRoutes(router *mux.Router, cfg config.Config, sr *service.Registry) {
	router.
		Methods("GET").
		Path("/internal/v1/sample").
		Handler(handlers.SampleHandlerV1())
}
