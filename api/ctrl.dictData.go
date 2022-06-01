package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// DictDataListAll godoc
// @Summary      字典数据列表查询
// @Description  获取所有字典数据信息
// @Tags         系统字典数据
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DictDataResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dictData/list [get]
func DictDataListAll(ctx *gin.Context) {
	res, err := service.DictDataList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// DictDataListTree godoc
// @Summary      字典数据列表查询
// @Description  获取所有字典数据信息
// @Tags         系统字典数据
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DictDataResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dictData/list [get]
func DictDataListTree(ctx *gin.Context) {
	res, err := service.DictDataList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// DictDataPage godoc
// @Summary      字典数据列表查询
// @Description  获取所有字典数据信息
// @Tags         系统字典数据
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DictDataResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dictData/p [post]
func DictDataPage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.DictDataPage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// DictDataCreate godoc
// @Summary      字典数据新增
// @Description  新增字典数据信息
// @Tags         系统字典数据
// @Accept       json
// @Produce      json
// @Param        req  body     request.DictDataCreateRequest  true  "新增字典数据信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dictData [post]
func DictDataCreate(ctx *gin.Context) {
	var req request.DictDataCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.DictDataCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// DictDataUpdate godoc
// @Summary      字典数据更新
// @Description  更新字典数据信息
// @Tags         系统字典数据
// @Accept       json
// @Produce      json
// @Param        req  body     request.DictDataUpdateRequest  true  "更新字典数据信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dictData [put]
func DictDataUpdate(ctx *gin.Context) {
	var req request.DictDataUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.DictDataUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// DictDataDelete godoc
// @Summary      字典数据伪删除
// @Description  伪删除字典数据信息
// @Tags         系统字典数据
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dictData/:id [delete]
func DictDataDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.DictDataDeleteFake(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// DictDataDetail godoc
// @Summary      字典数据详情
// @Description  详情字典数据信息
// @Tags         系统字典数据
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=response.DictDataResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dictData/:id [get]
func DictDataDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.DictDataDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
