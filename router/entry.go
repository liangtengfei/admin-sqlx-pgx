package router

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/router/business"
)

func entryRouter(root *gin.RouterGroup) {
	business.ArticleRouter(root)
	business.FileRouter(root)
}
