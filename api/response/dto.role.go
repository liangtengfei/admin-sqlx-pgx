package response

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// RoleResponse 角色信息
type RoleResponse struct {
	ID            int64          `json:"id"`                 // 角色ID
	RoleName      string         `json:"roleName"`           // 角色名称
	RoleKey       string         `json:"roleKey"`            // 角色权限字符串
	OrderNum      int32          `json:"orderNum"`           // 显示顺序
	DataScope     string         `json:"dataScope"`          // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status        string         `json:"status"`             // 帐号状态（0正常 1停用）
	DelFlag       string         `json:"delFlag"`            // 删除标志（0代表存在 2代表删除）
	CreateBy      string         `json:"createBy"`           // 创建者
	CreateTimeStr string         `json:"createTime"`         // 创建时间
	UpdateBy      string         `json:"updateBy"`           // 更新者
	UpdateTimeStr string         `json:"updateTime"`         // 更新时间
	Remark        string         `json:"remark"`             // 备注
	MenuList      []MenuResponse `json:"menuList,omitempty"` // 菜单列表
	DeptList      []DeptResponse `json:"deptList,omitempty"` // 部门列表
}

func (res *RoleResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *RoleResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}
