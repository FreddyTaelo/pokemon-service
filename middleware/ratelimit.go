package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(20, 100)

// 100 request per 5 seconds with 100 bursts
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded. Try again later.", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
