package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func UserRouter(r *gin.RouterGroup) {
	group := r.Group("user").
		Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.POST("p", api.UserPage)
		group.POST("", api.UserCreate)
		group.PUT("", api.UserUpdate)
		group.DELETE(":id", api.UserDelete)
		group.GET(":id", api.UserDetail)
		group.GET("detail/:id", api.UserDetailAll)
	}
}
