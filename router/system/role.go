package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	middleware2 "study.com/demo-sqlx-pgx/pkg/middleware"
)

func RoleRouter(r *gin.RouterGroup) {
	group := r.Group("role").Use(middleware2.AuthMiddleware(global.TokenMaker)).
		Use(middleware2.AccessControl(global.Enforcer))
	{
		group.GET("list", api.RoleListAll)
		group.POST("p", api.RolePage)
		group.POST("", api.RoleCreate)
		group.PUT("", api.RoleUpdate)
		group.DELETE(":id", api.RoleDelete)
		group.GET(":id", api.RoleDetail)
	}
}
