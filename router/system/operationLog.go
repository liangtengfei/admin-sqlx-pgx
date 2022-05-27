package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func OperationLogRouter(r *gin.RouterGroup) {
	group := r.Group("operlog").
		Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.POST("p", api.OperationLogPage)
	}
}
