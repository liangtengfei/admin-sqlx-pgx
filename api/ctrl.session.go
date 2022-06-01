package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// SessionPage godoc
// @Summary      操作日志列表查询
// @Description  获取所有操作日志信息
// @Tags         系统操作日志
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.SessionResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /post/p [post]
func SessionPage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.SessionPage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}
