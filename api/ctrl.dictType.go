package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// DictTypeListAll godoc
// @Summary      字典类型列表查询
// @Description  获取所有字典类型信息
// @Tags         系统字典类型
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DictTypeResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dictType/list [get]
func DictTypeListAll(ctx *gin.Context) {
	res, err := service.DictTypeList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// DictTypeListTree godoc
// @Summary      字典类型列表查询
// @Description  获取所有字典类型信息
// @Tags         系统字典类型
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DictTypeResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dictType/list [get]
func DictTypeListTree(ctx *gin.Context) {
	res, err := service.DictTypeList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// DictTypePage godoc
// @Summary      字典类型列表查询
// @Description  获取所有字典类型信息
// @Tags         系统字典类型
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DictTypeResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dictType/p [post]
func DictTypePage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.DictTypePage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// DictTypeCreate godoc
// @Summary      字典类型新增
// @Description  新增字典类型信息
// @Tags         系统字典类型
// @Accept       json
// @Produce      json
// @Param        req  body     request.DictTypeCreateRequest  true  "新增字典类型信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dictType [post]
func DictTypeCreate(ctx *gin.Context) {
	var req request.DictTypeCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.DictTypeCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// DictTypeUpdate godoc
// @Summary      字典类型更新
// @Description  更新字典类型信息
// @Tags         系统字典类型
// @Accept       json
// @Produce      json
// @Param        req  body     request.DictTypeUpdateRequest  true  "更新字典类型信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dictType [put]
func DictTypeUpdate(ctx *gin.Context) {
	var req request.DictTypeUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.DictTypeUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// DictTypeDelete godoc
// @Summary      字典类型伪删除
// @Description  伪删除字典类型信息
// @Tags         系统字典类型
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dictType/:id [delete]
func DictTypeDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.DictTypeDeleteFake(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// DictTypeDetail godoc
// @Summary      字典类型详情
// @Description  详情字典类型信息
// @Tags         系统字典类型
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=response.DictTypeResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dictType/:id [get]
func DictTypeDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.DictTypeDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
