package model

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// MenuResponse 菜单信息
type MenuResponse struct {
	ID            int64           `json:"id"`            // 菜单ID
	MenuName      string          `json:"menuName"`      // 菜单名称
	MenuKey       string          `json:"menuKey"`       // 菜单标识
	ParentID      int64           `json:"parentID"`      // 父菜单ID
	OrderNum      int32           `json:"orderNum"`      // 显示顺序
	Path          string          `json:"path"`          // 路由地址
	Component     string          `json:"component"`     // 组件路径
	RequestMethod string          `json:"requestMethod"` // 请求方法
	IsFrame       int32           `json:"isFrame"`       // 是否为外链（0是 1否）
	IsCache       int32           `json:"isCache"`       // 是否缓存（0缓存 1不缓存）
	MenuType      string          `json:"menuType"`      // 菜单类型（D目录 M菜单 A按钮）
	Visible       string          `json:"visible"`       // 菜单状态（0显示 1隐藏）
	Icon          string          `json:"icon"`          // 菜单图标
	Status        string          `json:"status"`        // 状态（0正常 1停用）
	DelFlag       string          `json:"delFlag"`       // 删除标志（0代表存在 2代表删除）
	CreateBy      string          `json:"createBy"`      // 创建者
	CreateTimeStr string          `json:"createTime"`    // 创建时间
	UpdateBy      string          `json:"updateBy"`      // 更新者
	UpdateTimeStr string          `json:"updateTime"`    // 更新时间
	Remark        string          `json:"remark"`        // 备注
	Children      []*MenuResponse `json:"children"`      // 子菜单
}

func (res *MenuResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *MenuResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}
