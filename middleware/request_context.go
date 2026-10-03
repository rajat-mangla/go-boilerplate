package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajat-mangla/go-boilerplate/apicontext"
)

func RequestContext() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqCtx := apicontext.NewRequestContextFromHTTP(r)
			ctx := apicontext.NewContextWithRequestContext(r.Context(), reqCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
