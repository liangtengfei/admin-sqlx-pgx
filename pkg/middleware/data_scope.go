package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"study.com/demo-sqlx-pgx/api/request"
	model "study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global/consts"
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
		var payload *token.Payload
		ctxGet, ok := ctx.Get(consts.AuthorizationPayloadKey)
		if ok {
			payload = ctxGet.(*token.Payload)
		} else {
			fail(ctx, http.StatusUnauthorized, "未查询到登录信息!")
			return
		}

		user, err := service.UserFindByUsername(ctx, payload.Username)
		if err != nil {
			fail(ctx, http.StatusUnauthorized, "未查询到用户信息!")
			return
		}

		var req request.DataScopeRequest

		// 超级管理员无需过滤
		if user.IsAdmin() {
			req = request.DataScopeRequest{
				Scope:  consts.ScopeDataAll.String(),
				Params: nil,
			}
			ctx.Set(consts.ScopeDataKey, req)
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

		req = request.DataScopeRequest{
			Scope:  role.DataScope,
			Params: scopeData,
		}

		// 通过上下文传递
		ctx.Set(consts.ScopeDataKey, req)

		ctx.Next()
	}
}
