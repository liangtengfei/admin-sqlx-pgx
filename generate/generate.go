package generate

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

const DSN = "postgresql://postgres:xiaohuozhi@localhost:5432/postgres?sslmode=disable"

const (
	TableCatalog = "postgres"
	TableSchema  = "public"
)

func GetDBConn() *pgxpool.Pool {
	//配置数据库
	ctx := context.Background()
	pgxConfig, err := pgxpool.ParseConfig(DSN)
	if err != nil {
		log.Fatal("配置pgx异常：", err)
	}

	conn, err := pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		log.Fatal("连接pgx异常：", err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal("联通数据库异常：", err)
	}
	return conn
}
