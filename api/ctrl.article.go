package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// ArticleListAll godoc
// @Summary      通知公告列表查询
// @Description  获取所有通知公告信息
// @Tags         通知公告
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=model.ArticleResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /article/list [get]
func ArticleListAll(ctx *gin.Context) {
	res, err := service.ArticleList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// ArticlePage godoc
// @Summary      通知公告分页查询
// @Description  分页获取所有通知公告信息
// @Tags         通知公告
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=model.ArticleResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /article/p [post]
func ArticlePage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.ArticlePage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// ArticleCreate godoc
// @Summary      通知公告新增
// @Description  通知公告新增信息
// @Tags         通知公告
// @Accept       json
// @Produce      json
// @Param        req  body     request.ArticleCreateRequest  true  "新增通知公告信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /article [post]
func ArticleCreate(ctx *gin.Context) {
	var req request.ArticleCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.ArticleCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// ArticleUpdate godoc
// @Summary      通知公告更新
// @Description  通知公告更新信息
// @Tags         通知公告
// @Accept       json
// @Produce      json
// @Param        req  body     request.ArticleUpdateRequest  true  "更新通知公告信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /article [put]
func ArticleUpdate(ctx *gin.Context) {
	var req request.ArticleUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.ArticleUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// ArticleDelete godoc
// @Summary      通知公告伪删除
// @Description  通知公告信息伪删除
// @Tags         通知公告
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /article/:id [delete]
func ArticleDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.ArticleDeleteFake(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// ArticleDetail godoc
// @Summary      通知公告详情
// @Description  通知公告详情信息
// @Tags         通知公告
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=model.ArticleResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /article/:id [get]
func ArticleDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.ArticleDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
