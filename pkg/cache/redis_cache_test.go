package cache

import (
	"context"
	v8 "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"study.com/demo-sqlx-pgx/config"
	"testing"
	"time"
)

func TestRedisCache(t *testing.T) {
	cache := NewRedisCache(config.Config{
		Redis: config.RedisConfig{
			Addr:     "127.0.0.1:6379",
			DB:       2,
			Password: "",
		}})
	ctx := context.Background()

	key := "CACHE:TEST"
	filed := "GO"
	value := "HELLO"

	err := cache.StrSet(ctx, key, value)
	require.NoError(t, err)

	getVal, err := cache.StrGet(ctx, key)
	require.NoError(t, err)
	require.Equal(t, value, getVal)

	err = cache.ExpireByKey(ctx, key, time.Second*1)
	require.NoError(t, err)

	time.Sleep(2 * time.Second)

	_, err = cache.StrGet(ctx, key)
	require.ErrorIs(t, err, v8.Nil)

	err = cache.HSet(ctx, key, filed, value)
	require.NoError(t, err)

	getValH, err := cache.HGet(ctx, key, filed)
	require.NoError(t, err)
	require.Equal(t, value, getValH)

	err = cache.ExpireByKey(ctx, key, time.Second*1)
	require.NoError(t, err)

	time.Sleep(2 * time.Second)

	values := map[string]interface{}{
		"A": "a",
		"B": "b",
	}
	err = cache.HMSet(ctx, key, values)
	require.NoError(t, err)

	hmgetVal, err := cache.HMGet(ctx, key)
	require.NoError(t, err)
	for k, v := range hmgetVal {
		require.Equal(t, values[k], v)
	}
}
