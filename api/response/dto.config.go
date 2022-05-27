package response

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// SysConfigResponse 参数配置表
type SysConfigResponse struct {
	ID            int32  `json:"id"`          // 参数主键
	ConfigName    string `json:"configName"`  // 参数名称
	ConfigKey     string `json:"configKey"`   // 参数键名
	ConfigValue   string `json:"configValue"` // 参数键值
	Status        string `json:"status"`      // 状态（0正常 1停用）
	DelFlag       string `json:"delFlag"`     // 删除标志（0代表存在 2代表删除）
	CreateBy      string `json:"createBy"`    // 创建者
	CreateTimeStr string `json:"createTime"`  // 创建时间
	UpdateBy      string `json:"updateBy"`    // 更新者
	UpdateTimeStr string `json:"updateTime"`  // 更新时间
	Remark        string `json:"remark"`      // 备注
}

func (res *SysConfigResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *SysConfigResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}
