package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func MenuRouter(r *gin.RouterGroup) {
	group := r.Group("menu").Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.GET("tree", api.MenuListTree)
		group.GET("list", api.MenuListAll)
		group.POST("p", api.MenuPage)
		group.POST("", api.MenuCreate)
		group.PUT("", api.MenuUpdate)
		group.DELETE(":id", api.MenuDelete)
		group.GET(":id", api.MenuDetail)
	}
}
