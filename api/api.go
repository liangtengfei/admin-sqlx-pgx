package api

import (
	"github.com/gin-gonic/gin"
	"study.com/demo-sqlx-pgx/api/response"
	model "study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global/consts"
	"study.com/demo-sqlx-pgx/pkg/token"
	"study.com/demo-sqlx-pgx/service"
)

func getAuthPayload(ctx *gin.Context) *token.Payload {
	return ctx.MustGet(consts.AuthorizationPayloadKey).(*token.Payload)
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
