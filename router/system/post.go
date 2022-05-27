package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	middleware2 "study.com/demo-sqlx-pgx/pkg/middleware"
)

func PostRouter(r *gin.RouterGroup) {
	group := r.Group("post").Use(middleware2.AuthMiddleware(global.TokenMaker)).
		Use(middleware2.AccessControl(global.Enforcer))
	{
		group.GET("list", api.PostListAll)
		group.POST("p", api.PostPage)
		group.POST("", api.PostCreate)
		group.PUT("", api.PostUpdate)
		group.DELETE(":id", api.PostDelete)
		group.GET(":id", api.PostDetail)
	}
}
