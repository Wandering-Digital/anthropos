package cache

import (
	"context"
	"time"
)

type Cache interface {
	BuildKey(keywords ...string) string
	Ping(ctx context.Context) error
	Set(ctx context.Context, key string, value string, expiration ...time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	Delete(ctx context.Context, key string) error
	IncrBy(ctx context.Context, key string, value int64) error
	DecrBy(ctx context.Context, key string, value int64) error
	Exists(ctx context.Context, key string) (int64, error)
	LRange(ctx context.Context, key string, start int, stop ...int) ([]string, error)
	LPush(ctx context.Context, key string, value string) error
}
