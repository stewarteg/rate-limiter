package tests

import (
	"testing"
	"time"

	"rate-limiter/limiter"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	rl := limiter.NewRateLimiter(client)

	key := "test-ip"
	limit := 5
	blockTime := 10 * time.Second

	for i := 0; i < limit; i++ {
		assert.True(t, rl.AllowRequest(key, limit, blockTime))
	}

	assert.False(t, rl.AllowRequest(key, limit, blockTime))
}
