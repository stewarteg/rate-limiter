package limiter

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(host, port string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
}
