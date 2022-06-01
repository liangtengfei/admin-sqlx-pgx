package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// UserPage godoc
// @Summary      用户分页查询
// @Description  分页获取所有用户信息
// @Tags         系统用户
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.RestRes{data=response.UserResponse}
// @Failure      500  {object}  response.RestRes
// @Router       /user/p [get]
func UserPage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.UserPageAndKeyword(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// UserCreate godoc
// @Summary      创建用户
// @Description  新增用户信息
// @Tags         系统用户
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.RestRes
// @Failure      500  {object}  response.RestRes
// @Router       /user [post]
func UserCreate(ctx *gin.Context) {
	var req request.UserCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	id, err := service.UserCreate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, id)
}

//query
//path
//header
//body
//formData

// UserUpdate godoc
// @Summary      更新用户
// @Description  更新用户信息
// @Tags         系统用户
// @Accept       json
// @Produce      json
// @Param        req  body     request.UserUpdateRequest  true  "更新用户信息"
// @Success      200  {object}  response.RestRes
// @Failure      500  {object}  response.RestRes
// @Router       /user [put]
func UserUpdate(ctx *gin.Context) {
	var req request.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.UserUpdate(ctx, req, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Success(ctx)
}

// UserDelete godoc
// @Summary      删除用户
// @Description  删除用户信息
// @Tags         系统用户
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200  {object}  response.RestRes
// @Failure      500  {object}  response.RestRes
// @Router       /user/:id      [delete]
func UserDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.UserDeleteById(ctx, req.Id, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, "删除失败："+err.Error())
		return
	}
	response.Success(ctx)
}

// UserDetail godoc
// @Summary      用户详情
// @Description  用户信息详情
// @Tags         系统用户
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200  {object}  response.RestRes
// @Failure      500  {object}  response.RestRes
// @Router       /user/:id      [get]
func UserDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	result, err := service.UserDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, result)
}

// UserDetailAll godoc
// @Summary      用户详情（关联部门、角色、岗位）
// @Description  用户信息详情（关联部门、角色、岗位）
// @Tags         系统用户
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200  {object}  response.RestRes
// @Failure      500  {object}  response.RestRes
// @Router       /user/:id      [get]
func UserDetailAll(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	result, err := service.UserDetail2(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, result)
}
