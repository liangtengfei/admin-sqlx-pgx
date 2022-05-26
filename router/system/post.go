package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func PostRouter(r *gin.RouterGroup) {
	group := r.Group("post").Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.GET("list", api.PostListAll)
		group.POST("p", api.PostPage)
		group.POST("", api.PostCreate)
		group.PUT("", api.PostUpdate)
		group.DELETE(":id", api.PostDelete)
		group.GET(":id", api.PostDetail)
	}
}
