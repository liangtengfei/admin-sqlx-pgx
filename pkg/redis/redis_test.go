package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"study.com/demo-sqlx-pgx/config"

	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
	cfg := config.Config{Redis: config.RedisConfig{
		Addr:     "127.0.0.1:6379",
		DB:       1,
		Password: "",
	}}
	rdb := NewRedis(cfg)

	ctx := context.Background()

	status := rdb.Ping(ctx)
	require.NoError(t, status.Err())

	rdb.Set(ctx, "hello", "redis", time.Second*100)

	get, err := rdb.Get(ctx, "hello").Result()
	require.NoError(t, err)
	require.Equal(t, "redis", get)
}

type Model struct {
	Str1    string   `redis:"str1"`
	Str2    string   `redis:"str2"`
	Int     int      `redis:"int"`
	Bool    bool     `redis:"bool"`
	Ignored struct{} `redis:"-"`
}

func TestScanToStruct(t *testing.T) {
	cfg := config.Config{Redis: config.RedisConfig{
		Addr:     "127.0.0.1:6379",
		DB:       1,
		Password: "",
	}}
	rdb := NewRedis(cfg)

	ctx := context.Background()
	_, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, "key", "str1", "hello")
		rdb.HSet(ctx, "key", "str2", "world")
		rdb.HSet(ctx, "key", "int", 123)
		rdb.HSet(ctx, "key", "bool", 1)
		return nil
	})
	require.NoError(t, err)

	var mode1 Model
	err = rdb.HGetAll(ctx, "key").Scan(&mode1)
	require.NoError(t, err)

	require.Equal(t, true, mode1.Bool)
}
