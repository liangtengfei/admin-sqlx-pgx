package request

// DictTypeCreateRequest 字典类型新增信息
type DictTypeCreateRequest struct {
	DictName string `json:"dictName" binding:"required"`          // 字典类型名称
	DictType string `json:"dictType" binding:"required,alphanum"` // 字典类型编码
	Remark   string `json:"remark" binding:"-"`                   // 备注
}

// DictTypeUpdateRequest 字典类型更新信息
type DictTypeUpdateRequest struct {
	ID       int64  `json:"id" binding:"required"`                // 字典主键
	DictName string `json:"dictName" binding:"required"`          // 字典类型名称
	DictType string `json:"dictType" binding:"required,alphanum"` // 字典类型编码
	Status   string `json:"status" binding:"-"`                   // 状态（0正常 1停用）
	Remark   string `json:"remark" binding:"-"`                   // 备注
}

// DictDataCreateRequest 字典数据新增信息
type DictDataCreateRequest struct {
	OrderNum  int32  `json:"orderNum" binding:"number,min=0"`      // 字典排序
	DictLabel string `json:"dictLabel" binding:"required"`         // 字典标签
	DictValue string `json:"dictValue" binding:"required"`         // 字典键值
	DictType  string `json:"dictType" binding:"required,alphanum"` // 字典类型编码
	CssClass  string `json:"cssClass" binding:"-"`                 // 样式属性（其他样式扩展）
	ListClass string `json:"listClass" binding:"-"`                // 表格回显样式
	Remark    string `json:"remark" binding:"-"`                   // 备注
}

// DictDataUpdateRequest 字典数据更新信息
type DictDataUpdateRequest struct {
	ID        int64  `json:"id" binding:"required"`           // 字典编码
	OrderNum  int32  `json:"orderNum" binding:"number,min=0"` // 字典排序
	DictLabel string `json:"dictLabel" binding:"required"`    // 字典标签
	DictValue string `json:"dictValue" binding:"required"`    // 字典键值
	CssClass  string `json:"cssClass" binding:"-"`            // 样式属性（其他样式扩展）
	ListClass string `json:"listClass" binding:"-"`           // 表格回显样式
	Status    string `json:"status" binding:"-"`              // 状态（0正常 1停用）
	Remark    string `json:"remark" binding:"-"`              // 备注
}
