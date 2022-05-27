package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	middleware2 "study.com/demo-sqlx-pgx/pkg/middleware"
)

func UserRouter(r *gin.RouterGroup) {
	group := r.Group("user").
		Use(middleware2.AuthMiddleware(global.TokenMaker)).
		Use(middleware2.AccessControl(global.Enforcer))
	{
		group.POST("p", api.UserPage)
		group.POST("", api.UserCreate)
		group.PUT("", api.UserUpdate)
		group.DELETE(":id", api.UserDelete)
		group.GET(":id", api.UserDetail)
		group.GET("detail/:id", api.UserDetailAll)
	}
}
