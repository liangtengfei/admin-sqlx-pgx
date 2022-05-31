package generate

import (
	"bufio"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func File2Lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func InsertStringToFile(path, content string, index int) error {
	lines, err := File2Lines(path)
	if err != nil {
		return err
	}

	var builder strings.Builder
	for i, line := range lines {
		if i == index {
			builder.WriteString(content)
		}
		builder.WriteString(line)
		builder.WriteString("\n")
	}

	return ioutil.WriteFile(path, []byte(builder.String()), 0666)
}

// InsertStringToFileEnd 插入末尾第几行 offset 跳过几行
func InsertStringToFileEnd(path, content string, offset int) error {
	lines, err := File2Lines(path)
	if err != nil {
		return err
	}

	index := len(lines) - offset

	var builder strings.Builder
	for i, line := range lines {
		if i == index {
			builder.WriteString(content)
			builder.WriteString("\n")
		}
		builder.WriteString(line)
		builder.WriteString("\n")
	}

	return ioutil.WriteFile(path, []byte(builder.String()), 0666)
}
