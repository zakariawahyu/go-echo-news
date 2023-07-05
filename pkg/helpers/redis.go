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

func KeyRedisType(basePrefix string, types string) string {
	return fmt.Sprintf("v2-%s-%s:", basePrefix, types)
}

func KeyRedisTypeKey(basePrefix string, types string, key string, limit int, offset int) string {
	if types != "" && limit != 0 {
		return fmt.Sprintf("v2-%s-%s-%s:%d-%d", basePrefix, types, key, limit, offset)
	} else if limit == 0 {
		return fmt.Sprintf("v2-%s-%s-%s:", basePrefix, types, key)
	}
	return fmt.Sprintf("v2-%s:%d-%d", basePrefix, limit, offset)
}

func KeyRedisTypeFeatured(basePrefix string, types string, featured bool, limit int, offset int) string {
	if featured {
		return fmt.Sprintf("v2-%s-featured-%s:%d-%d", basePrefix, types, limit, offset)
	}
	return fmt.Sprintf("v2-%s-%s:%d-%d", basePrefix, types, limit, offset)
}

func KeyRedisMultimediaTypeKey(multimediaType string, basePrefix string, types string, key string, limit int, offset int) string {
	if types == "" {
		return fmt.Sprintf("v2-%s-%s:%d-%d", multimediaType, basePrefix, limit, offset)
	}

	return fmt.Sprintf("v2-%s-%s-%s-%s:%d-%d", multimediaType, basePrefix, types, key, limit, offset)
}

func KeyRedisTypeKeyDate(basePrefix string, types string, key string, date string, limit int, offset int) string {
	if types == "" {
		return fmt.Sprintf("v2-%s:%d-%d", basePrefix, limit, offset)
	}

	return fmt.Sprintf("v2-%s-%s-%s-%s:%d-%d", basePrefix, types, key, date, limit, offset)
}

func TtlRedis(ttl int) time.Duration {
	return time.Second * time.Duration(ttl)
}
