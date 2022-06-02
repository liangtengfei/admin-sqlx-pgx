package redis

import (
	v8 "github.com/go-redis/redis/v8"
	"study.com/demo-sqlx-pgx/config"
)

func NewRedis(config config.Config) *v8.Client {
	return v8.NewClient(&v8.Options{
		Addr:     config.Redis.Addr,
		DB:       config.Redis.DB,
		Password: config.Redis.Password,
	})
}
