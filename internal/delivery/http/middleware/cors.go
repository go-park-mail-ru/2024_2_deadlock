package middleware

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
)

func CorsMW(cfg *bootstrap.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			for _, v := range cfg.Server.CorsAllowOrigins {
				if v == origin {
					w.Header().Set("Access-Control-Allow-Origin", v)
					w.Header().Set("Access-Control-Allow-Credentials", "true")

					break
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
