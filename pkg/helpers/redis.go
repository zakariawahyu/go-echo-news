package helpers

import (
	"fmt"
	"time"
)

var (
	Fastest = 60      // 1 minute
	Faster  = 120     // 2 minute
	Fast    = 1800    // 30 minute
	Slow    = 3600    // 60 minute
	Slowly  = 10800   // 3 hours
	HalfDay = 43200   // 12 hours
	Slower  = 86400   // 24 hours
	Slowest = 86400   // 3 days
	Long    = 604800  // 7 days
	Longer  = 1209600 // 14 days
	Longest = 2592000 // 30 days
)

func KeyRedis(basePrefix string, key string) string {
	return fmt.Sprintf("v2-%s:%s", basePrefix, key)
}

func KeyRedisRowContent(basePrefix string, types string, key string, limit int, offset int) string {
	if types != "" {
		return fmt.Sprintf("v2-%s-%s-%s:%d-%d", basePrefix, types, key, limit, offset)
	}
	return fmt.Sprintf("v2-%s:%d-%d", basePrefix, limit, offset)
}

func TtlRedis(ttl int) time.Duration {
	return time.Second * time.Duration(ttl)
}
