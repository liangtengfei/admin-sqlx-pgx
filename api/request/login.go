package request

// LoginRequest 登录信息
type LoginRequest struct {
	Username      string `json:"username" binding:"required"`              // 登录名称
	Password      string `json:"password" binding:"required,min=6,max=30"` // 登录密码
	CaptchaId     string `json:"id" binding:"-"`                           // 验证码
	CaptchaAnswer string `json:"code" binding:"-"`                         // 验证码编号
}

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
