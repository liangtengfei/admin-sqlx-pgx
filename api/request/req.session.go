package request

import (
	"github.com/google/uuid"
	"time"
)

type SessionCreateRequest struct {
	ID           uuid.UUID `db:"id"`            // 唯一标识
	UserName     string    `db:"user_name"`     // 用户名
	RealName     string    `db:"real_name"`     // 真实姓名
	RefreshToken string    `db:"refresh_token"` // 刷新秘钥
	UserAgent    string    `db:"user_agent"`    // 请求信息
	ClientIp     string    `db:"client_ip"`     // 请求地址
	IsBlocked    bool      `db:"is_blocked"`    // 是否阻止
	ExpiresAt    time.Time `db:"expires_at"`    // 过期时间
	CreateAt     time.Time `db:"create_a"`      // 创建时间
	Remark       string    `db:"remark"`        // 备注
}

type SessionUpdateRequest struct {
	ID        uuid.UUID `db:"id"`         // 唯一标识
	IsBlocked bool      `db:"is_blocked"` // 是否阻止
	Remark    string    `db:"remark"`     // 备注
}
