package request

// SysConfigCreateRequest 配置新增信息
type SysConfigCreateRequest struct {
	ConfigName  string `json:"configName" binding:"required"`  // 参数名称
	ConfigKey   string `json:"configKey" binding:"required"`   // 参数键名
	ConfigValue string `json:"configValue" binding:"required"` // 参数键值
	Remark      string `json:"remark" binding:"-"`             // 备注
}

// SysConfigUpdateRequest 配置更新信息
type SysConfigUpdateRequest struct {
	ID          int64  `json:"id" binding:"required"`          // 参数主键
	ConfigName  string `json:"configName" binding:"required"`  // 参数名称
	ConfigValue string `json:"configValue" binding:"required"` // 参数键值
	Status      string `json:"status" binding:"-"`             // 状态（0正常 1停用）
	Remark      string `json:"remark" binding:"-"`             // 备注
}
