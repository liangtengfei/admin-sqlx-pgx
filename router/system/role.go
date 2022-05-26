package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func RoleRouter(r *gin.RouterGroup) {
	group := r.Group("role").Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.GET("list", api.RoleListAll)
		group.POST("p", api.RolePage)
		group.POST("", api.RoleCreate)
		group.PUT("", api.RoleUpdate)
		group.DELETE(":id", api.RoleDelete)
		group.GET(":id", api.RoleDetail)
	}
}
