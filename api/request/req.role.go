package request

// RoleCreateRequest 角色新增信息
type RoleCreateRequest struct {
	RoleName  string  `json:"roleName" binding:"required"`         // 角色名称
	RoleKey   string  `json:"roleKey" binding:"required,alphanum"` // 角色权限字符串
	OrderNum  int32   `json:"orderNum" binding:"number,min=0"`     // 显示顺序
	DataScope string  `json:"dataScope" binding:"required"`        // 数据范围
	Remark    string  `json:"remark" binding:"-"`                  // 备注
	DeptIds   []int64 `json:"deptIds" binding:"-"`                 // 部门编号数组
	MenuIds   []int64 `json:"menuIds" binding:"-"`                 // 菜单编号数据
}

// RoleUpdateRequest 角色更新信息
type RoleUpdateRequest struct {
	ID        int64   `json:"id" binding:"required"`               // 角色ID
	RoleName  string  `json:"roleName" binding:"required"`         // 角色名称
	RoleKey   string  `json:"roleKey" binding:"required,alphanum"` // 角色权限字符串
	OrderNum  int32   `json:"orderNum" binding:"number,min=0"`     // 显示顺序
	DataScope string  `json:"dataScope" binding:"required"`        // 数据范围
	Status    string  `json:"status" binding:"-"`                  // 状态（0正常 1停用）
	Remark    string  `json:"remark" binding:"-"`                  // 备注
	DeptIds   []int64 `json:"deptIds" binding:"-"`                 // 部门编号数组
	MenuIds   []int64 `json:"menuIds" binding:"-"`                 // 菜单编号数据
}
