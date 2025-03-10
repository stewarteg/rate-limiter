package limiter

type LimiterStrategy interface {
	AllowRequest(key string, limit int, blockTime int) bool
}
