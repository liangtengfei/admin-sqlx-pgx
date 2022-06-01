package main

import (
	"embed"
	"fmt"
	"study.com/demo-sqlx-pgx/bootstrap"
)

//go:embed resources/public/*
var static embed.FS

// @title                       Admin GOGOGO
// @version                     0.0.1
// @description                 This is a sample golang project
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath                    127.0.0.1:5000
func main() {
	fmt.Println("HI! GO + SQLX + PGX")

	bootstrap.RunServer(static)
}
