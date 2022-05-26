package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// SysConfigListAll godoc
// @Summary      参数配置列表查询
// @Description  获取所有参数配置信息
// @Tags         系统参数配置
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=model.SysConfigResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /config/list [get]
func SysConfigListAll(ctx *gin.Context) {
	res, err := service.ConfigList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// SysConfigPage godoc
// @Summary      参数配置列表查询
// @Description  获取所有参数配置信息
// @Tags         系统参数配置
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=model.SysConfigResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /config/p [post]
func SysConfigPage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.ConfigPage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// SysConfigCreate godoc
// @Summary      参数配置新增
// @Description  新增参数配置信息
// @Tags         系统参数配置
// @Accept       json
// @Produce      json
// @Param        req  body     request.SysConfigCreateRequest  true  "新增参数配置信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /config [post]
func SysConfigCreate(ctx *gin.Context) {
	var req request.SysConfigCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.ConfigCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// SysConfigUpdate godoc
// @Summary      参数配置更新
// @Description  更新参数配置信息
// @Tags         系统参数配置
// @Accept       json
// @Produce      json
// @Param        req  body     request.SysConfigUpdateRequest  true  "更新参数配置信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /config [put]
func SysConfigUpdate(ctx *gin.Context) {
	var req request.SysConfigUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.ConfigUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// SysConfigDelete godoc
// @Summary      参数配置伪删除
// @Description  伪删除参数配置信息
// @Tags         系统参数配置
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /config/:id [delete]
func SysConfigDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.ConfigDeleteFake(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// SysConfigDetail godoc
// @Summary      参数配置详情
// @Description  详情参数配置信息
// @Tags         系统参数配置
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=model.SysConfigResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /config/:id [get]
func SysConfigDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.ConfigDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}

func SysConfigBatchInsert(ctx *gin.Context) {
	var req []request.SysConfigCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	log.Println(len(req), req)
	res, err := service.ConfigBatchCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
