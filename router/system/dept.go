package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func DeptRouter(r *gin.RouterGroup) {
	// 需要数据过滤的路由
	scopeGroup := r.Group("dept").Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer)).
		Use(middleware.DataScope())
	{
		scopeGroup.POST("p", api.DeptPage)
	}

	group := r.Group("dept").
		Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.GET("list", api.DeptListAll)
		group.GET("tree", api.DeptListTree)
		group.POST("", api.DeptCreate)
		group.PUT("", api.DeptUpdate)
		group.DELETE(":id", api.DeptDelete)
		group.GET(":id", api.DeptDetail)
	}
}
