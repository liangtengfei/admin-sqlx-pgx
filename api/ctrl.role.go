package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// RoleListAll godoc
// @Summary      角色列表查询
// @Description  获取所有角色信息
// @Tags         系统角色
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=model.RoleResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /role/list [get]
func RoleListAll(ctx *gin.Context) {
	res, err := service.RoleList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// RolePage godoc
// @Summary      角色列表查询
// @Description  获取所有角色信息
// @Tags         系统角色
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=model.RoleResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /role/p [post]
func RolePage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.RolePage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// RoleCreate godoc
// @Summary      角色新增
// @Description  新增角色信息
// @Tags         系统角色
// @Accept       json
// @Produce      json
// @Param        req  body     request.RoleCreateRequest  true  "新增角色信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /role [post]
func RoleCreate(ctx *gin.Context) {
	var req request.RoleCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.RoleCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// RoleUpdate godoc
// @Summary      角色更新
// @Description  更新角色信息
// @Tags         系统角色
// @Accept       json
// @Produce      json
// @Param        req  body     request.RoleUpdateRequest  true  "更新角色信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /role [put]
func RoleUpdate(ctx *gin.Context) {
	var req request.RoleUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.RoleUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// RoleDelete godoc
// @Summary      角色伪删除
// @Description  伪删除角色信息
// @Tags         系统角色
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /role/:id [delete]
func RoleDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.RoleDeleteFake(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// RoleDetail godoc
// @Summary      角色详情
// @Description  详情角色信息
// @Tags         系统角色
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=model.RoleResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /role/:id [get]
func RoleDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.RoleDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

func RoleListByUserId(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	res, err := service.RoleListByUserId(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}
