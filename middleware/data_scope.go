package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/global/consts"
	"study.com/demo-sqlx-pgx/model"
	"study.com/demo-sqlx-pgx/pkg/token"
	"study.com/demo-sqlx-pgx/service"
)

func DataScope() gin.HandlerFunc {
	fail := func(ctx *gin.Context, code int, msg string) {
		ctx.AbortWithStatusJSON(code, gin.H{
			"message": msg,
			"code":    code,
		})
	}
	return func(ctx *gin.Context) {
		payload := ctx.MustGet(consts.AuthorizationPayloadKey).(*token.Payload)

		user, err := service.UserFindByUsername(ctx, payload.Username)
		if err != nil {
			fail(ctx, http.StatusUnauthorized, "未查询到用户信息!")
			return
		}

		// 超级管理员无需过滤
		if user.IsAdmin() {
			ctx.Next()
			return
		}

		if len(user.RoleList) == 0 {
			fail(ctx, http.StatusUnauthorized, "请配置角色信息")
			return
		}

		//获取最高权限
		roleList := make([]model.RoleResponse, len(user.RoleList))
		copy(roleList, user.RoleList)
		sort.Slice(roleList, func(i, j int) bool {
			return roleList[i].DataScope < roleList[j].DataScope
		})
		role := roleList[0]

		var scopeData []interface{}
		if role.DataScope == consts.ScopeDataAll.String() {
		} else if role.DataScope == consts.ScopeCustom.String() {
			scopeData = append(scopeData, role.ID)
		} else if role.DataScope == consts.ScopeDept.String() {
			scopeData = append(scopeData, user.DeptID)
		} else if role.DataScope == consts.ScopeDeptChild.String() {
			scopeData = append(scopeData, user.DeptID)
		} else {
			scopeData = append(scopeData, user.UserName)
		}

		req := request.DataScopeRequest{
			Scope:  role.DataScope,
			Params: scopeData,
		}

		// 通过上下文传递
		ctx.Set(consts.ScopeDataKey, req)

		ctx.Next()
	}
}
