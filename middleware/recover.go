package middleware

import (
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// RecoverMiddleware catches panics and prevents the request goroutine from crashing.
func RecoverMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					path, _ := mux.CurrentRoute(r).GetPathTemplate()
					if path == "" {
						path = r.URL.String()
					}

					path = NewNormalizedPath(path)

					log.Error().Msgf("Recovered from panic: %+v\npath: %s\n%s", err, path, string(debug.Stack()))

					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// NewNormalizedPath cleans up request path and request router matching path to just keep
// alphanumeric chars and remove other charaters.
// Some systems like statsd, new relic etc
// like clean slugified like string for monitoring.
func NewNormalizedPath(path string) string {
	if strings.ContainsRune(path, ':') {
		replaceRe := regexp.MustCompile(`\{(.+?)(:.+?)(\{[0-9-]+\})?\}`)
		path = replaceRe.ReplaceAllString(path, "{$1}")
	}

	return path
}
