package response

import (
	"database/sql"
	"github.com/google/uuid"
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// UserResponse 用户信息表
type UserResponse struct {
	ID            int64          `json:"id"`                 // 用户ID
	DeptID        int64          `json:"deptID"`             // 部门ID
	UserName      string         `json:"userName"`           // 用户账号
	RealName      string         `json:"realName"`           // 用户昵称
	Email         string         `json:"email"`              // 用户邮箱
	Mobile        string         `json:"mobile"`             // 手机号码
	Sex           string         `json:"sex"`                // 用户性别（0男 1女 2未知）
	Avatar        string         `json:"avatar"`             // 头像地址
	Password      string         `json:"-"`                  // 密码
	Status        string         `json:"status"`             // 帐号状态（0正常 1停用）
	DelFlag       string         `json:"delFlag"`            // 删除标志（0代表存在 2代表删除）
	CreateBy      string         `json:"createBy"`           // 创建者
	CreateTimeStr string         `json:"createTime"`         // 创建时间
	UpdateBy      string         `json:"updateBy"`           // 更新者
	UpdateTimeStr string         `json:"updateTime"`         // 更新时间
	Remark        string         `json:"remark"`             // 备注
	DeptNameStr   string         `json:"deptName,omitempty"` // 部门名称
	RoleKeys      []string       `json:"roles,omitempty"`    // 角色信息
	Dept          DeptResponse   `json:"dept,omitempty"`     // 部门信息
	RoleList      []RoleResponse `json:"roleList,omitempty"` // 角色信息
	PostList      []PostResponse `json:"postList,omitempty"` // 岗位信息

}

func (res *UserResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *UserResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}

func (res *UserResponse) DeptName(deptName sql.NullString) {
	if deptName.Valid {
		res.DeptNameStr = deptName.String
	} else {
		res.DeptNameStr = ""
	}
}

func (res *UserResponse) IsAdmin() bool {
	is := false
	if len(res.RoleList) > 0 {
		for _, role := range res.RoleList {
			if role.RoleKey == "admin" {
				is = true
				break
			}
		}
	}
	if res.ID == int64(1) {
		is = true
	}
	return is
}

type LoginResponse struct {
	SessionID             uuid.UUID    `json:"sessionID"`
	AccessToken           string       `json:"accessToken"`
	AccessTokenExpiresAt  time.Time    `json:"accessTokenExpiresAt"`
	RefreshToken          string       `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time    `json:"refreshTokenExpiresAt"`
	User                  UserResponse `json:"user"`
	Roles                 []string     `json:"roles"`
}
