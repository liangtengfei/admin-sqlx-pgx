package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"study.com/demo-sqlx-pgx/global/consts"
	"study.com/demo-sqlx-pgx/pkg/token"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(consts.AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("未提供授权头部")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg":  err.Error(),
				"code": http.StatusUnauthorized,
			})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("授权头部格式不正确")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg":  err.Error(),
				"code": http.StatusUnauthorized,
			})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != consts.AuthorizationTypeBearer {
			err := fmt.Errorf("不支持的授权头部类型 %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg":  err.Error(),
				"code": http.StatusUnauthorized,
			})
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg":  err.Error(),
				"code": http.StatusUnauthorized,
			})
			return
		}

		ctx.Set(consts.AuthorizationUsernameKey, payload.Username)
		ctx.Set(consts.AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
