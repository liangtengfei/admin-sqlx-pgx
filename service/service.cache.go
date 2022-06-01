package service

import (
	"context"
	"study.com/demo-sqlx-pgx/pkg/redis"
)

func CacheConfigInRedis() error {
	ctx := context.Background()

	cfgs, err := store.ConfigList(ctx)
	if err != nil {
		return err
	}

	values := map[string]interface{}{}
	for _, cfg := range cfgs {
		values[cfg.ConfigKey] = cfg.ConfigValue
	}

	return redis.ConfigCache(ctx, values)
}

func CacheConfigByKey(key string) (string, error) {
	ctx := context.Background()

	return redis.ConfigByKey(ctx, key)
}
