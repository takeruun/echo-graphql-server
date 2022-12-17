package middlewares

import (
	"app/helpers"
	"context"
	"net/http"
)

func SetHttpContextMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpContext := helpers.HTTP{
				W: &w,
				R: r,
			}
			httpKeyContext := helpers.HTTPKey("http")

			// context 経由で http.ResponseWriter, *http.Request の値使えるようにする
			ctx := context.WithValue(r.Context(), httpKeyContext, httpContext)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
