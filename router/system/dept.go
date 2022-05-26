package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func DeptRouter(r *gin.RouterGroup) {
	group := r.Group("dept").Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.GET("list", api.DeptListAll)
		group.GET("tree", api.DeptListTree)
		group.POST("p", api.DeptPage)
		group.POST("", api.DeptCreate)
		group.PUT("", api.DeptUpdate)
		group.DELETE(":id", api.DeptDelete)
		group.GET(":id", api.DeptDetail)
	}
}
