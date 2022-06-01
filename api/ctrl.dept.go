package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// DeptListAll godoc
// @Summary      部门列表查询
// @Description  获取所有部门信息
// @Tags         系统部门
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DeptResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dept/list [get]
func DeptListAll(ctx *gin.Context) {
	res, err := service.DeptList(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// DeptListTree godoc
// @Summary      部门列表查询(树形）
// @Description  获取所有部门信息(树形）
// @Tags         系统部门
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DeptResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dept/tree [get]
func DeptListTree(ctx *gin.Context) {
	res, err := service.DeptListTree(ctx)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, res)
}

// DeptPage godoc
// @Summary      部门列表查询
// @Description  获取所有部门信息
// @Tags         系统部门
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.DeptResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dept/p [post]
func DeptPage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.DeptPage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// DeptCreate godoc
// @Summary      部门新增
// @Description  新增部门信息
// @Tags         系统部门
// @Accept       json
// @Produce      json
// @Param        req  body     request.DeptCreateRequest  true  "新增部门信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dept [post]
func DeptCreate(ctx *gin.Context) {
	var req request.DeptCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.DeptCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// DeptUpdate godoc
// @Summary      部门更新
// @Description  更新部门信息
// @Tags         系统部门
// @Accept       json
// @Produce      json
// @Param        req  body     request.DeptUpdateRequest  true  "更新部门信息"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dept [put]
func DeptUpdate(ctx *gin.Context) {
	var req request.DeptUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	err := service.DeptUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// DeptDelete godoc
// @Summary      部门伪删除
// @Description  伪删除部门信息
// @Tags         系统部门
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /dept/:id [delete]
func DeptDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.DeptDeleteFake(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// DeptDetail godoc
// @Summary      部门详情
// @Description  详情部门信息
// @Tags         系统部门
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=response.DeptResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /dept/:id [get]
func DeptDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.DeptDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
