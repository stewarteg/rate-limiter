package middleware

import (
	"net/http"
	"rate-limiter/limiter"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var rateLimiter *limiter.RateLimiter

func init() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost" + ":" + "6379",
	})
	rateLimiter = limiter.NewRateLimiter(client)
}

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		token := r.Header.Get("API_KEY")

		limit, _ := strconv.Atoi("5")
		blockTime, _ := strconv.Atoi("300")

		if token != "" {
			limit, _ = strconv.Atoi("10")
		}

		if !rateLimiter.AllowRequest(ip, limit, time.Duration(blockTime)*time.Second) {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
