package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/middleware"
)

func NoticeRouter(r *gin.RouterGroup) {
	group := r.Group("notice").Use(middleware.AuthMiddleware(global.TokenMaker)).
		Use(middleware.AccessControl(global.Enforcer))
	{
		group.POST("p", api.NoticePage)
		group.POST("", api.NoticeCreate)
		group.PUT("", api.NoticeUpdate)
		group.DELETE(":id", api.NoticeDelete)
		group.GET(":id", api.NoticeDetail)
	}
}
