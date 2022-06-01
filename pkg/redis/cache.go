package redis

import (
	"context"
	"errors"
)

const (
	SysConfigKey = "SYS:CONFIG"
)

func ConfigCache(ctx context.Context, values map[string]interface{}) error {
	if len(values) <= 0 {
		return nil
	}
	return Conn.HSet(ctx, SysConfigKey, values).Err()
}

func ConfigByKey(ctx context.Context, field string) (string, error) {
	if len(field) <= 0 {
		return "", errors.New("KEY不能为空")
	}

	return Conn.HGet(ctx, SysConfigKey, field).Result()
}
