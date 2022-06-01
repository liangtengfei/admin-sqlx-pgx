package redis

import (
	"github.com/go-redis/redis/v8"
	"study.com/demo-sqlx-pgx/config"
)

func NewRedis(config config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		DB:       config.Redis.DB,
		Password: config.Redis.Password,
	})
}

func InitRedis(config config.Config) {
	Conn = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		DB:       config.Redis.DB,
		Password: config.Redis.Password,
	})
}

var Conn *redis.Client
