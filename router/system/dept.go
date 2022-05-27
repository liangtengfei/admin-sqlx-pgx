package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	middleware2 "study.com/demo-sqlx-pgx/pkg/middleware"
)

func DeptRouter(r *gin.RouterGroup) {
	// 需要数据过滤的路由
	scopeGroup := r.Group("dept").Use(middleware2.AuthMiddleware(global.TokenMaker)).
		Use(middleware2.AccessControl(global.Enforcer)).
		Use(middleware2.DataScope())
	{
		scopeGroup.POST("p", api.DeptPage)
	}

	group := r.Group("dept").
		Use(middleware2.AuthMiddleware(global.TokenMaker)).
		Use(middleware2.AccessControl(global.Enforcer))
	{
		group.GET("list", api.DeptListAll)
		group.GET("tree", api.DeptListTree)
		group.POST("", api.DeptCreate)
		group.PUT("", api.DeptUpdate)
		group.DELETE(":id", api.DeptDelete)
		group.GET(":id", api.DeptDetail)
	}
}
