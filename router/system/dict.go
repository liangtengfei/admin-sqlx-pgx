package system

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api"
	"study.com/demo-sqlx-pgx/global"
	middleware2 "study.com/demo-sqlx-pgx/pkg/middleware"
)

func DictRouter(r *gin.RouterGroup) {
	group := r.Group("dict").Use(middleware2.AuthMiddleware(global.TokenMaker)).
		Use(middleware2.AccessControl(global.Enforcer))
	{
		group.GET("type/list", api.DictTypeListAll)
		group.POST("type/p", api.DictTypePage)
		group.POST("type", api.DictTypeCreate)
		group.PUT("type", api.DictTypeUpdate)
		group.DELETE("type/:id", api.DictTypeDelete)
		group.GET("type/:id", api.DictTypeDetail)

		group.GET("data/list", api.DictDataListAll)
		group.POST("data/p", api.DictDataPage)
		group.POST("data", api.DictDataCreate)
		group.PUT("data", api.DictDataUpdate)
		group.DELETE("data/:id", api.DictDataDelete)
		group.GET("data/:id", api.DictDataDetail)
	}
}
