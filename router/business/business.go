package business

import "github.com/gin-gonic/gin"

func InitBusinessRouter(root *gin.RouterGroup) {
	ArticleRouter(root)
}
