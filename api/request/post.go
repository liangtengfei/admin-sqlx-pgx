package request

// PostCreateRequest 岗位新增信息
type PostCreateRequest struct {
	PostName string `json:"postName" binding:"required"`     // 岗位名称
	OrderNum int32  `json:"orderNum" binding:"number,min=0"` // 显示顺序
	Remark   string `json:"remark" binding:"-"`              // 备注
}

// PostUpdateRequest 岗位更新信息
type PostUpdateRequest struct {
	ID       int64  `json:"id" binding:"required"`           // 岗位ID
	PostName string `json:"postName" binding:"required"`     // 岗位名称
	OrderNum int32  `json:"orderNum" binding:"number,min=0"` // 显示顺序
	Status   string `json:"status" binding:"-"`              // 状态（0正常 1停用）
	Remark   string `json:"remark" binding:"-"`              // 备注
}
