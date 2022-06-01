package business

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/pkg/middleware"
)

// FileRouter 路由配置
func FileRouter(r *gin.RouterGroup) {
	group := r.Group("file").
		Use(middleware.AuthMiddleware(global.TokenMaker))
	{
		group.POST("upload/common", api.FileUploadCommon)
		group.POST("p", api.FilePage)
		group.DELETE(":id", api.FileDelete)
		group.GET(":id", api.FileDetail)
	}
}
