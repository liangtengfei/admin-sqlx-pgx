package model

import (
	"github.com/google/uuid"
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

type SessionResponse struct {
	ID           uuid.UUID `json:"id"`           // 唯一标识
	UserName     string    `json:"userName"`     // 用户名
	RealName     string    `json:"realName"`     // 真实姓名
	RefreshToken string    `json:"refreshToken"` // 刷新秘钥
	UserAgent    string    `json:"userAgent"`    // 请求信息
	ClientIp     string    `json:"clientIp"`     // 请求地址
	IsBlocked    bool      `json:"isBlocked"`    // 是否阻止
	ExpiresAt    time.Time `json:"expiresAt"`    // 是否阻止
	CreateAtStr  string    `json:"createAt"`     // 创建时间
	Remark       string    `json:"remark"`       // 备注
}

func (res *SessionResponse) CreateAt(createAt time.Time) {
	res.CreateAtStr = datetime.ToDatetime(createAt)
}
