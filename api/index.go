package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/response"
)

// Index godoc
// @Summary      首页
// @Description  展示简单信息
// @Tags         首页
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.RestRes
// @Failure      500  {object}  response.RestRes
// @Router       / [get]
func Index(ctx *gin.Context) {
	response.SuccessData(ctx, gin.H{"PING": "PONG"})
}
