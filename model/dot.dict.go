package model

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// DictTypeResponse 字典类型表
type DictTypeResponse struct {
	ID            int64              `json:"id"`                 // 字典主键
	DictName      string             `json:"dictName"`           // 字典类型名称
	DictType      string             `json:"dictType"`           // 字典类型编码
	Status        string             `json:"status"`             // 状态（0正常 1停用）
	DelFlag       string             `json:"delFlag"`            // 删除标志（0代表存在 2代表删除）
	CreateBy      string             `json:"createBy"`           // 创建者
	CreateTimeStr string             `json:"createTime"`         // 创建时间
	UpdateBy      string             `json:"updateBy"`           // 更新者
	UpdateTimeStr string             `json:"updateTime"`         // 更新时间
	Remark        string             `json:"remark"`             // 备注,
	DataList      []DictDataResponse `json:"dataList,omitempty"` // 字典数据列表
}

func (res *DictTypeResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *DictTypeResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}

// DictDataResponse 字典数据表
type DictDataResponse struct {
	ID            int64  `json:"id"`         // 字典编码
	OrderNum      int32  `json:"orderNum"`   // 字典排序
	DictLabel     string `json:"dictLabel"`  // 字典标签
	DictValue     string `json:"dictValue"`  // 字典键值
	DictType      string `json:"dictType"`   // 字典类型编码
	CssClass      string `json:"cssClass"`   // 样式属性（其他样式扩展）
	ListClass     string `json:"listClass"`  // 表格回显样式
	Status        string `json:"status"`     // 状态（0正常 1停用）
	DelFlag       string `json:"delFlag"`    // 删除标志（0代表存在 2代表删除）
	CreateBy      string `json:"createBy"`   // 创建者
	CreateTimeStr string `json:"createTime"` // 创建时间
	UpdateBy      string `json:"updateBy"`   // 更新者
	UpdateTimeStr string `json:"updateTime"` // 更新时间
	Remark        string `json:"remark"`     // 备注
}

func (res *DictDataResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *DictDataResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}
