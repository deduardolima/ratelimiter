package middleware

import (
	"net/http"
	"strings"

	config "github.com/deduardolima/ratelimiter/configs"
	"github.com/deduardolima/ratelimiter/infra/database"
	"github.com/deduardolima/ratelimiter/internal/limiter"
	"github.com/gorilla/mux"
)

func RateLimiterMiddleware(store database.Store, cfg *config.Config) mux.MiddlewareFunc {
	rateLimiter := limiter.NewRateLimiter(store, cfg)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if strings.Contains(ip, ":") {
				ip = strings.Split(ip, ":")[0]
			}

			token := r.Header.Get("API_KEY")

			blockedIP, _ := rateLimiter.IsBlocked(ip)
			blockedToken, _ := rateLimiter.IsBlocked(token)

			if blockedIP || blockedToken {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}

			if !rateLimiter.AllowRequest(ip, token) {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
