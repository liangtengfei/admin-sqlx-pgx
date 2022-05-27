package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/model"
	"study.com/demo-sqlx-pgx/service"
	"study.com/demo-sqlx-pgx/utils"
	"study.com/demo-sqlx-pgx/utils/captcha"
	"time"
)

func Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	//验证验证码
	if req.CaptchaId != "" && req.CaptchaAnswer != "" {
		err := captcha.VerifyCaptcha(req.CaptchaId, req.CaptchaAnswer)
		if err != nil {
			response.ErrorMsg(ctx, err.Error())
			return
		}
	}

	// 用户名或者手机号码都能登录
	user, err := service.UserFindByUsername(ctx, req.Username)
	if err != nil && err != sql.ErrNoRows {
		response.ErrorMsg(ctx, "用户查找失败")
		return
	}
	if err != nil && err == sql.ErrNoRows {
		user, err = service.UserFindByMobile(ctx, req.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				response.ErrorMsg(ctx, "用户不存在")
				return
			} else {
				response.ErrorMsg(ctx, "用户查找失败")
				return
			}
		}
	}

	err = utils.CheckPassword(req.Password+"_"+user.Mobile, user.Password)
	if err != nil {
		response.ErrorMsg(ctx, "密码不正确")
		return
	}

	//生成token
	accessToken, accessPayload, err := global.TokenMaker.CreateToken(req.Username, global.Config.Auth.AccessTokenDuration)
	if err != nil {
		response.ErrorMsg(ctx, "登录失败，生成token失败")
		return
	}
	refreshToken, refreshPayload, err := global.TokenMaker.CreateToken(req.Username, global.Config.Auth.RefreshTokenDuration)
	if err != nil {
		response.ErrorMsg(ctx, "登录失败，生成refresh token失败")
		return
	}

	sessionReq := request.SessionCreateRequest{
		ID:           refreshPayload.ID,
		UserName:     user.UserName,
		RealName:     user.RealName,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
		CreateAt:     time.Now(),
		Remark:       "",
	}

	// 插入记录
	go func() {
		id, err := service.SessionCreate(ctx, sessionReq, user.UserName)
		if err != nil {
			global.Log.Error("用户登录", zap.String("TAG", "新增会话记录"), zap.Error(err))
			return
		}
		global.Log.Info("用户登录", zap.String("TAG", id.String()))
	}()

	rsp := model.LoginResponse{
		SessionID:             accessPayload.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  user,
		Roles:                 []string{},
	}

	response.SuccessData(ctx, rsp)
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
}

func RefreshToken(ctx *gin.Context) {
	var req request.RenewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	refreshPayload, err := global.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":  err.Error(),
			"code": http.StatusUnauthorized,
		})
		return
	}

	session, err := service.SessionDetail(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			response.ErrorCodeMsg(ctx, http.StatusNotFound, "未查询到记录")
			return
		}
		response.ErrorCodeMsg(ctx, http.StatusInternalServerError, "查询记录失败")
		return
	}

	if session.IsBlocked {
		response.ErrorMsg(ctx, "该会话已经被阻止")
		return
	}

	if session.UserName != refreshPayload.Username {
		response.ErrorMsg(ctx, "用户信息不匹配")
		return
	}

	if time.Now().After(session.ExpiresAt) {
		response.ErrorMsg(ctx, "会话已过期")
		return
	}

	//生成token
	accessToken, accessPayload, err := global.TokenMaker.CreateToken(refreshPayload.Username, global.Config.Auth.AccessTokenDuration)
	if err != nil {
		response.ErrorMsg(ctx, "登录失败，生成token失败")
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	response.SuccessData(ctx, rsp)
}

func GenerateCaptcha(ctx *gin.Context) {
	res, err := captcha.GenerateCaptcha("string")
	if err != nil {
		response.ErrorMsg(ctx, "获取验证码错误")
		return
	}
	response.SuccessData(ctx, res)
}
