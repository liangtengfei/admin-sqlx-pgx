package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	middleware2 "study.com/demo-sqlx-pgx/pkg/middleware"
)

func OperationLogRouter(r *gin.RouterGroup) {
	group := r.Group("operlog").
		Use(middleware2.AuthMiddleware(global.TokenMaker)).
		Use(middleware2.AccessControl(global.Enforcer))
	{
		group.POST("p", api.OperationLogPage)
	}
}
