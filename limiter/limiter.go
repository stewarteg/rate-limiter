package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimiter struct {
	client *redis.Client
}

func NewRateLimiter(client *redis.Client) *RateLimiter {
	return &RateLimiter{client: client}
}

type LimiterStrategy interface {
	AllowRequest(key string, limit int, blockTime int) bool
}

func (rl *RateLimiter) AllowRequest(key string, limit int, blockTime time.Duration) bool {
	ctx := context.Background()
	pipe := rl.client.TxPipeline()

	// contador
	count := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Second)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false
	}

	// Verificar se passa do limite
	if count.Val() > int64(limit) {
		rl.client.Set(ctx, key+":blocked", "1", blockTime)
		return false
	}

	return true
}
