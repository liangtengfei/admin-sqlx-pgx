package main

import (
	"fmt"
	"study.com/demo-sqlx-pgx/bootstrap"
)

// @title                       Admin GOGOGO
// @version                     0.0.1
// @description                 This is a sample golang project
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath                    127.0.0.1:5000
func main() {
	fmt.Println("HI! GO + SQLX + PGX")

	bootstrap.RunServer()
}
