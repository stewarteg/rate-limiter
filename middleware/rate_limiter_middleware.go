package middleware

import (
	"errors"
	"net/http"
	"rate-limiter/limiter"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

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
			requestLimit, err := decodeJWT(token)
			if err == nil {
				limit = requestLimit
			}
		}

		if !rateLimiter.AllowRequest(ip, limit, time.Duration(blockTime)*time.Second) {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func decodeJWT(tokenString string) (int, error) {
	type CustomClaims struct {
		RequestLimit string `json:"request_limit"`
		jwt.RegisteredClaims
	}

	token, _, err := jwt.NewParser().ParseUnverified(tokenString, &CustomClaims{})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Converte o valor de "request_limit" para int
	requestLimit, err := strconv.Atoi(claims.RequestLimit)
	if err != nil {
		return 0, err
	}

	return requestLimit, nil
}
