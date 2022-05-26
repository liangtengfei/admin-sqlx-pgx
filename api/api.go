package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global/consts"
	"study.com/demo-sqlx-pgx/model"
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
