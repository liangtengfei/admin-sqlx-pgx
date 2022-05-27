package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sort"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	model "study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global/consts"
	"study.com/demo-sqlx-pgx/pkg/token"
	"study.com/demo-sqlx-pgx/service"
)

func getAuthPayload(ctx *gin.Context) *token.Payload {
	return ctx.MustGet(consts.AuthorizationPayloadKey).(*token.Payload)
}

func GetSessionId(ctx *gin.Context) uuid.UUID {
	return getAuthPayload(ctx).ID
}

func GetLoginUserName(ctx *gin.Context) string {
	return getAuthPayload(ctx).Username
}

func GetLoginUserInfo(ctx *gin.Context) model.UserResponse {
	info, err := service.UserFindByUsername(ctx, getAuthPayload(ctx).Username)
	if err != nil {
		response.ErrorMsg(ctx, "未查询到登录信息，请重新登录")
		return info
	}
	return info
}

func DataScopeRequest(ctx *gin.Context) {
	user := GetLoginUserInfo(ctx)

	if len(user.RoleList) == 0 {
		//err = errors.New("请配置角色信息")
		response.ErrorCodeMsg(ctx, http.StatusUnauthorized, "请配置角色信息")
		return
	}

	var scope string
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

	// 管理员全部数据
	if role.RoleKey == "admin" {
		scope = consts.ScopeDataAll.String()
	} else {
		scope = role.DataScope
	}

	req := request.DataScopeRequest{
		Scope:  scope,
		Params: scopeData,
	}

	// 通过上下文传递
	ctx.Set(consts.ScopeDataKey, req)
	return
}
