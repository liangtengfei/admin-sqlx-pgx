package request

// MenuCreateRequest 菜单信息 新增
type MenuCreateRequest struct {
	MenuName      string `json:"menuName" binding:"required"`                                // 菜单名称
	MenuKey       string `json:"menuKey" binding:"required"`                                 // 菜单标识
	ParentID      int64  `json:"parentID" binding:"number,min=0"`                            // 父菜单ID
	OrderNum      int32  `json:"orderNum" binding:"number,min=0"`                            // 显示顺序
	Path          string `json:"path" binding:"required"`                                    // 路由地址
	Component     string `json:"component" binding:"-"`                                      // 组件路径
	IsFrame       string `json:"isFrame" binding:"-"`                                        // 是否为外链（0是 1否）
	IsCache       string `json:"isCache" binding:"-"`                                        // 是否缓存（0缓存 1不缓存）
	MenuType      string `json:"menuType" binding:"oneof=D M A"`                             // 菜单类型（D目录 M菜单 A按钮）
	Visible       string `json:"visible" binding:"-"`                                        // 菜单状态（0显示 1隐藏）
	Icon          string `json:"icon" binding:"-"`                                           // 菜单图标
	Remark        string `json:"remark" binding:"-"`                                         // 备注
	RequestMethod string `json:"requestMethod" binding:"required,oneof=GET POST DELETE PUT"` // 请求方法
}

// MenuUpdateRequest 菜单信息 更新
type MenuUpdateRequest struct {
	ID            int64  `json:"id" binding:"required"`                                      // 菜单ID
	MenuName      string `json:"menuName" binding:"required"`                                // 菜单名称
	MenuKey       string `json:"menuKey" binding:"required"`                                 // 菜单标识
	ParentID      int64  `json:"parentID" binding:"number,min=0"`                            // 父菜单ID
	OrderNum      int32  `json:"orderNum" binding:"number,min=0"`                            // 显示顺序
	Path          string `json:"path" binding:"required"`                                    // 路由地址
	Component     string `json:"component" binding:"-"`                                      // 组件路径
	IsFrame       string `json:"isFrame" binding:"-"`                                        // 是否为外链（0是 1否）
	IsCache       string `json:"isCache" binding:"-"`                                        // 是否缓存（0缓存 1不缓存）
	MenuType      string `json:"menuType" binding:"oneof=D M A"`                             // 菜单类型（D目录 M菜单 A按钮）
	Visible       string `json:"visible" binding:"-"`                                        // 菜单状态（0显示 1隐藏）
	Icon          string `json:"icon" binding:"-"`                                           // 菜单图标
	Status        string `json:"status" binding:"-"`                                         // 状态（0正常 1停用）
	Remark        string `json:"remark" binding:"-"`                                         // 备注
	RequestMethod string `json:"requestMethod" binding:"required,oneof=GET POST DELETE PUT"` // 请求方法
}
