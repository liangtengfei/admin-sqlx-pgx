package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// MenuListTree godoc
// @Summary      菜单列表树
// @Description  获取所有菜单信息
// @Tags         系统菜单
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.MenuResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /menu/tree [get]
func MenuListTree(ctx *gin.Context) {
	res, err := service.MenuListTree(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}

// MenuListAll godoc
// @Summary      岗位列表查询
// @Description  获取所有岗位信息
// @Tags         系统菜单
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.MenuResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /menu/list [get]
func MenuListAll(ctx *gin.Context) {
	res, err := service.MenuListTree(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// MenuPage godoc
// @Summary      岗位列表查询
// @Description  获取所有岗位信息
// @Tags         系统菜单
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.MenuResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /menu/p [post]
func MenuPage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.MenuPage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// MenuCreate godoc
// @Summary      岗位新增
// @Description  新增岗位信息
// @Tags         系统菜单
// @Accept       json
// @Produce      json
// @Param        req  body     request.MenuCreateRequest  true  "新增岗位信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /menu [post]
func MenuCreate(ctx *gin.Context) {
	var req request.MenuCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.MenuCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// MenuUpdate godoc
// @Summary      岗位更新
// @Description  更新岗位信息
// @Tags         系统菜单
// @Accept       json
// @Produce      json
// @Param        req  body     request.MenuUpdateRequest  true  "更新岗位信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /menu [put]
func MenuUpdate(ctx *gin.Context) {
	var req request.MenuUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.MenuUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// MenuDelete godoc
// @Summary      岗位伪删除
// @Description  伪删除岗位信息
// @Tags         系统菜单
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /menu/:id [delete]
func MenuDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.MenuDeleteFake(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// MenuDetail godoc
// @Summary      岗位详情
// @Description  详情岗位信息
// @Tags         系统菜单
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=response.MenuResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /menu/:id [get]
func MenuDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.MenuDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
