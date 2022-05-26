package request

// DeptCreateRequest 部门新增信息
type DeptCreateRequest struct {
	ParentID  int64  `json:"parentID" binding:"number,min=0"` // 父部门id
	DeptName  string `json:"deptName" binding:"required"`     // 部门名称
	DeptCode  string `json:"deptCode" binding:"required"`     // 部门编码
	OrderNum  int32  `json:"orderNum" binding:"number,min=0"` // 显示顺序
	Leader    string `json:"leader" binding:"-"`              // 负责人
	Phone     string `json:"phone" binding:"-"`               // 联系电话
	Email     string `json:"email" binding:"-"`               // 邮箱
	Remark    string `json:"remark" binding:"-"`              // 备注
	Ancestors string `binding:"-"`                            // 祖籍
}

// DeptUpdateRequest 部门更新信息
type DeptUpdateRequest struct {
	ID        int64  `json:"id" binding:"required"`           // 部门id
	ParentID  int64  `json:"parentID" binding:"number,min=0"` // 父部门id
	DeptName  string `json:"deptName" binding:"required"`     // 部门名称
	DeptCode  string `json:"deptCode" binding:"required"`     // 部门编码
	OrderNum  int32  `json:"orderNum" binding:"number,min=0"` // 显示顺序
	Leader    string `json:"leader" binding:"-"`              // 负责人
	Phone     string `json:"phone" binding:"-"`               // 联系电话
	Email     string `json:"email" binding:"-"`               // 邮箱
	Status    string `json:"status" binding:"required"`       // 状态（0正常 1停用）
	Remark    string `json:"remark" binding:"-"`              // 备注
	Ancestors string `binding:"-"`                            // 祖籍
}
