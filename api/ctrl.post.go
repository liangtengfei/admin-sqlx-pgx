package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// PostListAll godoc
// @Summary      岗位列表查询
// @Description  获取所有岗位信息
// @Tags         系统岗位
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.PostResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /post/list [get]
func PostListAll(ctx *gin.Context) {
	res, err := service.PostList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// PostPage godoc
// @Summary      岗位列表查询
// @Description  获取所有岗位信息
// @Tags         系统岗位
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.PostResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /post/p [post]
func PostPage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.PostPage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// PostCreate godoc
// @Summary      岗位新增
// @Description  新增岗位信息
// @Tags         系统岗位
// @Accept       json
// @Produce      json
// @Param        req  body     request.PostCreateRequest  true  "新增岗位信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /post [post]
func PostCreate(ctx *gin.Context) {
	var req request.PostCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.PostCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// PostUpdate godoc
// @Summary      岗位更新
// @Description  更新岗位信息
// @Tags         系统岗位
// @Accept       json
// @Produce      json
// @Param        req  body     request.PostUpdateRequest  true  "更新岗位信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /post [put]
func PostUpdate(ctx *gin.Context) {
	var req request.PostUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.PostUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// PostDelete godoc
// @Summary      岗位伪删除
// @Description  伪删除岗位信息
// @Tags         系统岗位
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /post/:id [delete]
func PostDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.PostDeleteFake(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// PostDetail godoc
// @Summary      岗位详情
// @Description  详情岗位信息
// @Tags         系统岗位
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=response.PostResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /post/:id [get]
func PostDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.PostDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
