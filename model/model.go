package model

// BaseModel 基础model
type BaseModel struct {
	DelFlag    string `json:"-"`          // 删除标记Y是N否
	Status     int    `json:"status"`     // 状态 0默认
	Remark     string `json:"remark"`     // 备注
	CreateBy   string `json:"createBy"`   // 创建人员
	CreateTime string `json:"createTime"` // 创建时间
	UpdateBy   string `json:"updateBy"`   // 修改人员
	UpdateTime string `json:"updateTime"` // 修改时间
}
