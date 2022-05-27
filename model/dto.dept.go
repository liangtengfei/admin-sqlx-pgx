package model

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// DeptResponse 部门信息
type DeptResponse struct {
	ID            int64           `json:"id"`                     // 部门id
	ParentID      int64           `json:"parentID"`               // 父部门id
	Ancestors     string          `json:"ancestors"`              // 祖级列表
	DeptName      string          `json:"deptName"`               // 部门名称
	DeptCode      string          `json:"deptCode"`               // 部门编码
	OrderNum      int32           `json:"orderNum"`               // 显示顺序
	Status        string          `json:"status"`                 // 状态（0正常 1停用）
	DelFlag       string          `json:"delFlag"`                // 删除标志（0代表存在 2代表删除）
	CreateBy      string          `json:"createBy"`               // 创建者
	CreateTimeStr string          `json:"createTime"`             // 创建时间
	UpdateBy      string          `json:"updateBy"`               // 更新者
	UpdateTimeStr string          `json:"updateTime"`             // 更新时间
	Remark        string          `json:"remark"`                 // 备注
	Children      []*DeptResponse `json:"children,omitempty"`     // 字部门列表
	ChildrenSize  int             `json:"childrenSize,omitempty"` // 子部门数量
}

func (res *DeptResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *DeptResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}
