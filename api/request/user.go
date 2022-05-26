package request

// UserCreateRequest 用户新增信息
type UserCreateRequest struct {
	UserName string  `json:"userName" binding:"required,alphanum"`                     // 用户账号
	RealName string  `json:"realName" binding:"required"`                              // 用户昵称
	Email    string  `json:"email" binding:"-"`                                        // 用户邮箱
	DeptID   int64   `json:"deptID" binding:"required"`                                // 部门ID
	Remark   string  `json:"remark" binding:"-"`                                       // 备注
	Password string  `json:"password" binding:"required" minLength:"6" maxLength:"16"` // 密码
	Avatar   string  `json:"avatar" binding:"-"`                                       // 头像地址
	Sex      string  `json:"sex" binding:"-" enums:"0, 1, 2"`                          // 用户性别（0男 1女 2未知）
	Mobile   string  `json:"mobile" binding:"required"`                                // 手机号码
	RoleIds  []int64 `json:"roleIds" binding:"-" example:"1,2,3"`                      // 角色编号数组
	PostIds  []int64 `json:"PostIds" binding:"-" example:"1,2,3"`                      // 岗位编号数组
}

// UserUpdateRequest 用户更新信息
type UserUpdateRequest struct {
	ID       int64   `json:"id" binding:"required"`                              // 用户ID
	DeptID   int64   `json:"deptID" binding:"required"`                          // 部门ID
	RealName string  `json:"realName" binding:"required"`                        // 用户昵称
	Mobile   string  `json:"mobile" binding:"required"`                          // 手机号码
	Email    string  `json:"email" binding:"required"`                           // 用户邮箱
	Sex      string  `json:"sex" binding:"required,oneof=0 1 2" enums:"0, 1, 2"` // 用户性别（0男 1女 2未知）
	Avatar   string  `json:"avatar" binding:"required"`                          // 头像地址
	Status   string  `json:"status" binding:"-"`                                 // 状态（0正常 1停用）
	Remark   string  `json:"remark" binding:"required"`                          // 备注
	RoleIds  []int64 `json:"roleIds" binding:"-" example:"1,2,3"`                // 角色编号数组
	PostIds  []int64 `json:"PostIds" binding:"-" example:"1,2,3"`                // 岗位编号数组
}
