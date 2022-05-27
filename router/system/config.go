package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func SysConfigRouter(r *gin.RouterGroup) {
	group := r.Group("config").Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.GET("list", api.SysConfigListAll)
		group.POST("p", api.SysConfigPage)
		group.POST("", api.SysConfigCreate)
		group.POST("batch", api.SysConfigBatchInsert)
		group.PUT("", api.SysConfigUpdate)
		group.DELETE(":id", api.SysConfigDelete)
		group.GET(":id", api.SysConfigDetail)
	}
}
