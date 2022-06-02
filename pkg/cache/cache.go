package cache

import (
	"context"
	"time"
)

type Store interface {
	StrSet(ctx context.Context, key string, values string) error
	StrGet(ctx context.Context, key string) (string, error)
	HSet(ctx context.Context, key string, field string, value interface{}) error
	HGet(ctx context.Context, key string, field string) (string, error)
	HMSet(ctx context.Context, key string, values map[string]interface{}) error
	HMGet(ctx context.Context, key string) (map[string]string, error)
	ExpireByKey(ctx context.Context, key string, duration time.Duration) error
}
