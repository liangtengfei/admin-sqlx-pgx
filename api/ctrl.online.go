package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/service"
)

// OnlineUserMenus godoc
// @Summary      用户菜单
// @Description  登录用户菜单信息
// @Tags         登录用户
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.RestRes
// @Failure      500  {object}  response.RestRes
// @Router       /user/online/userInfo      [get]
func OnlineUserMenus(ctx *gin.Context) {
	var result = make(map[string]interface{}, 0)
	user, err := service.UserOnlyByUsername(ctx, GetLoginUserName(ctx))
	if err != nil {
		response.ErrorMsg(ctx, "查找用户信息失败")
		return
	}

	if user.Posts != "" && len(user.Posts) > 0 {
		posts, err := service.PostListByIds(ctx, user.Posts)
		if err != nil {
			response.ErrorMsg(ctx, "查找岗位信息失败")
			return
		}
		result["posts"] = posts
	}

	if user.DeptID > 0 {
		dept, err := service.DeptDetail(ctx, user.DeptID)
		if err != nil {
			response.ErrorMsg(ctx, "查找部门信息失败")
			return
		}
		user.Dept = dept
	}

	roles, err := service.RoleListByUserId(ctx, user.ID)
	if err != nil {
		response.ErrorMsg(ctx, "查找角色信息失败")
		return
	}
	result["roles"] = roles

	if len(roles) > 0 {
		var roleIds []int64
		for _, role := range roles {
			roleIds = append(roleIds, role.ID)
		}
		menus, err := service.MenuListByRoleIds(ctx, roleIds)
		if err != nil {
			response.ErrorMsg(ctx, "查找菜单信息失败")
			return
		}
		result["menus"] = menus
	}
	result["user"] = user
	response.SuccessData(ctx, result)
}
