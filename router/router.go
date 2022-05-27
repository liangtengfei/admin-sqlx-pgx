package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/pkg/middleware"
	_ "study.com/demo-sqlx-pgx/resources/docs"
	"study.com/demo-sqlx-pgx/router/system"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//静态资源
	r.Static("/favicon.ico", "./resources/public/favicon.ico")
	r.Static("/static", "./resources/public/assets") // dist

	root := r.Group("/")
	root.Use(middleware.Cors())
	root.GET("", api.Index)
	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	root.POST("login", api.Login)
	root.POST("refreshToken", api.RefreshToken)
	root.GET("/captcha", api.GenerateCaptcha)

	system.UserRouter(root)
	system.DeptRouter(root)
	system.RoleRouter(root)
	system.MenuRouter(root)
	system.PostRouter(root)
	system.DictRouter(root)
	system.NoticeRouter(root)
	system.SysConfigRouter(root)
	system.OperationLogRouter(root)
	system.SessionRouter(root)

	return r
}
