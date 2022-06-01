package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// NoticePage godoc
// @Summary      通知公告列表查询
// @Description  获取所有通知公告信息
// @Tags         系统通知公告
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.NoticeResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /post/p [post]
func NoticePage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.NoticePage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// NoticeCreate godoc
// @Summary      通知公告新增
// @Description  新增通知公告信息
// @Tags         系统通知公告
// @Accept       json
// @Produce      json
// @Param        req  body     request.NoticeCreateRequest  true  "新增通知公告信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /post [post]
func NoticeCreate(ctx *gin.Context) {
	var req request.NoticeCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.NoticeCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// NoticeUpdate godoc
// @Summary      通知公告更新
// @Description  更新通知公告信息
// @Tags         系统通知公告
// @Accept       json
// @Produce      json
// @Param        req  body     request.NoticeUpdateRequest  true  "更新通知公告信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /post [put]
func NoticeUpdate(ctx *gin.Context) {
	var req request.NoticeUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.NoticeUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// NoticeDelete godoc
// @Summary      通知公告伪删除
// @Description  伪删除通知公告信息
// @Tags         系统通知公告
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /post/:id [delete]
func NoticeDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.NoticeDelete(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// NoticeDetail godoc
// @Summary      通知公告详情
// @Description  详情通知公告信息
// @Tags         系统通知公告
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=response.NoticeResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /post/:id [get]
func NoticeDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.NoticeDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
