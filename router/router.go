package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/fs"
	"net/http"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/pkg/middleware"
	_ "study.com/demo-sqlx-pgx/resources/docs"
	"study.com/demo-sqlx-pgx/router/system"
)

func InitRouter(staticFs embed.FS) *gin.Engine {
	r := gin.Default()

	//静态资源
	//r.Static("/favicon.ico", "./resources/public/favicon.ico")
	//r.Static("/static", "./resources/public/assets") // dist

	fe, _ := fs.Sub(staticFs, "resources/public")
	r.StaticFS("/static", http.FS(fe))

	root := r.Group("/")
	root.Use(middleware.Cors())
	root.GET("", api.Index)
	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	root.POST("login", api.Login)
	root.POST("refreshToken", api.RefreshToken)
	root.GET("/captcha", api.GenerateCaptcha)

	system.UserRouter(root)
	system.UserOnlineRouter(root)
	system.DeptRouter(root)
	system.RoleRouter(root)
	system.MenuRouter(root)
	system.PostRouter(root)
	system.DictRouter(root)
	system.NoticeRouter(root)
	system.SysConfigRouter(root)
	system.OperationLogRouter(root)
	system.SessionRouter(root)

	// 其他业务路由
	entryRouter(root)

	return r
}
