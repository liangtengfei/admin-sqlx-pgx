package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	middleware2 "study.com/demo-sqlx-pgx/pkg/middleware"
)

func SysConfigRouter(r *gin.RouterGroup) {
	group := r.Group("config").Use(middleware2.AuthMiddleware(global.TokenMaker)).
		Use(middleware2.AccessControl(global.Enforcer))
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
