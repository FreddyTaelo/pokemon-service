package middleware

import (
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"pokemon-service/config"
	"strconv"
)

var limiter *rate.Limiter

func init() {
	cfg := config.LoadConfig()

	rateLimit, err := strconv.ParseFloat(cfg.RateLimitPerSecond, 64)
	if err != nil {
		log.Fatalf("Invalid rate limit value: %v", err)
	}
	rateLimitBurst, err := strconv.Atoi(cfg.RateLimitBurst)
	if err != nil {
		log.Fatalf("Invalid rate limit value: %v", err)
	}

	limiter = rate.NewLimiter(rate.Limit(rateLimit), rateLimitBurst) // 100 requests burst
}

// 100 requests per 5 seconds with 100 bursts
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded (Too Many Requests). Try again later", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
} /*RateLimitMiddleware()*/
