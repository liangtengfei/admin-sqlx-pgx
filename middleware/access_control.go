package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"study.com/demo-sqlx-pgx/global/consts"
	"study.com/demo-sqlx-pgx/pkg/token"
	"study.com/demo-sqlx-pgx/service"
)

func AccessControl(enforcer *casbin.Enforcer) gin.HandlerFunc {
	fail := func(ctx *gin.Context, code int, msg string) {
		ctx.AbortWithStatusJSON(code, gin.H{
			"message": msg,
			"code":    code,
		})
	}
	return func(ctx *gin.Context) {
		obj := ctx.Request.URL.Path
		act := ctx.Request.Method

		val, exists := ctx.Get(consts.AuthorizationPayloadKey)
		if !exists {
			fail(ctx, http.StatusUnauthorized, "请先登录!")
			return
		}
		payload := val.(*token.Payload)

		user, err := service.UserFindByUsername(ctx, payload.Username)
		if err != nil {
			fail(ctx, http.StatusUnauthorized, "未查询到用户信息!")
			return
		}

		if len(user.RoleKeys) == 0 {
			fail(ctx, http.StatusUnauthorized, "无访问权限!")
			return
		}

		for _, role := range user.RoleKeys {
			ok, err := enforcer.Enforce(role, obj, act)
			if ok {
				ctx.Next()
				return
			}
			if err != nil || !ok {
				continue
			}
		}
		fail(ctx, http.StatusUnauthorized, "无访问权限!")
		return
	}
}
