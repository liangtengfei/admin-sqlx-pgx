package business

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/pkg/middleware"
)

// ArticleRouter 路由配置
func ArticleRouter(r *gin.RouterGroup) {
	group := r.Group("article").
		Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.POST("p", api.ArticlePage)
		group.POST("", api.ArticleCreate)
		group.PUT("", api.ArticleUpdate)
		group.DELETE(":id", api.ArticleDelete)
		group.GET(":id", api.ArticleDetail)
	}
}
