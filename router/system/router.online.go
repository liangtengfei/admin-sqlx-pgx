package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	middleware2 "study.com/demo-sqlx-pgx/pkg/middleware"
)

func UserOnlineRouter(r *gin.RouterGroup) {
	group := r.Group("user/online").
		Use(middleware2.AuthMiddleware(global.TokenMaker))
	{
		group.POST("userInfo", api.OnlineUserMenus)
	}
}
