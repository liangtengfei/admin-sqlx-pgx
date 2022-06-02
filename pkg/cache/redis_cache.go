package cache

import (
	"context"
	"errors"
	v8 "github.com/go-redis/redis/v8"
	"study.com/demo-sqlx-pgx/config"
	"study.com/demo-sqlx-pgx/pkg/redis"
	"time"
)

var NotFoundErr = errors.New("未查询到数据")

type RedisCache struct {
	client *v8.Client
}

func NewRedisCache(config config.Config) RedisCache {
	return RedisCache{
		client: redis.NewRedis(config),
	}
}

func (c RedisCache) StrSet(ctx context.Context, key string, value string) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

func (c RedisCache) StrGet(ctx context.Context, key string) (string, error) {
	res, err := c.client.Get(ctx, key).Result()
	if err == v8.Nil {
		return "", NotFoundErr
	}
	return res, nil
}

func (c RedisCache) HSet(ctx context.Context, key string, field string, value interface{}) error {
	values := map[string]interface{}{field: value}
	return c.client.HSet(ctx, key, values).Err()
}

func (c RedisCache) HGet(ctx context.Context, key string, field string) (string, error) {
	res, err := c.client.HGet(ctx, key, field).Result()
	if err == v8.Nil {
		return "", NotFoundErr
	}
	return res, nil
}

func (c RedisCache) HMSet(ctx context.Context, key string, values map[string]interface{}) error {
	return c.client.HSet(ctx, key, values).Err()
}

func (c RedisCache) HMGet(ctx context.Context, key string) (map[string]string, error) {
	res, err := c.client.HGetAll(ctx, key).Result()
	if err == v8.Nil {
		return nil, NotFoundErr
	}
	return res, nil
}

func (c RedisCache) ExpireByKey(ctx context.Context, key string, duration time.Duration) error {
	return c.client.Expire(ctx, key, duration).Err()
}
