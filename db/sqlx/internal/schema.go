package internal

import (
	"time"

	"github.com/google/uuid"
)

type AgoConfig struct {
	ID          int64     `db:"id"`           // 唯一标识
	ConfigName  string    `db:"config_name"`  // 配置名称
	ConfigKey   string    `db:"config_key"`   // 配置标识
	ConfigValue string    `db:"config_value"` // 配置内容
	Status      string    `db:"status"`       // 状态（默认0）
	DelFlag     string    `db:"del_flag"`     // 删除标记
	CreateTime  time.Time `db:"create_time"`  // 创建时间
	UpdateTime  time.Time `db:"update_time"`  // 更新时间
	CreateBy    string    `db:"create_by"`    // 创建人员
	UpdateBy    string    `db:"update_by"`    // 更新人员
	Remark      string    `db:"remark"`       // 备注
}

type AgoDept struct {
	ID         int64     `db:"id"`          // 唯一标识
	ParentID   int64     `db:"parent_id"`   // 父部门id
	DeptName   string    `db:"dept_name"`   // 部门名称
	DeptCode   string    `db:"dept_code"`   // 部门编码
	Ancestors  string    `db:"ancestors"`   // 祖级列表
	OrderNum   int32     `db:"order_num"`   // 显示顺序
	Status     string    `db:"status"`      // 状态（默认0）
	DelFlag    string    `db:"del_flag"`    // 删除标记
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateBy   string    `db:"create_by"`   // 创建人员
	UpdateBy   string    `db:"update_by"`   // 更新人员
	Remark     string    `db:"remark"`      // 备注
}

type AgoDictData struct {
	ID         int64     `db:"id"`          // 唯一标识
	DictLabel  string    `db:"dict_label"`  // 字典标签
	DictValue  string    `db:"dict_value"`  // 字典键值
	DictType   string    `db:"dict_type"`   // 字典类型编码
	ListClass  string    `db:"list_class"`  // 表格回显样式
	CssClass   string    `db:"css_class"`   // 样式属性（其他样式扩展）
	Status     string    `db:"status"`      // 状态（默认0）
	DelFlag    string    `db:"del_flag"`    // 删除标记
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateBy   string    `db:"create_by"`   // 创建人员
	UpdateBy   string    `db:"update_by"`   // 更新人员
	Remark     string    `db:"remark"`      // 备注
	OrderNum   int32     `db:"order_num"`   // 排序编号
}

type AgoDictType struct {
	ID         int64     `db:"id"`          // 字典主键
	DictName   string    `db:"dict_name"`   // 字典类型名称
	DictType   string    `db:"dict_type"`   // 字典类型编码
	Status     string    `db:"status"`      // 状态（默认0）
	DelFlag    string    `db:"del_flag"`    // 删除标记
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateBy   string    `db:"create_by"`   // 创建人员
	UpdateBy   string    `db:"update_by"`   // 更新人员
	Remark     string    `db:"remark"`      // 备注
}

type AgoMenu struct {
	ID         int64     `db:"id"`          // 唯一标识
	MenuName   string    `db:"menu_name"`   // 菜单名称
	MenuKey    string    `db:"menu_key"`    // 菜单标识
	ParentID   int64     `db:"parent_id"`   // 上级标识
	Path       string    `db:"path"`        // 路由地址
	MenuType   string    `db:"menu_type"`   // 菜单类型（D目录 M菜单 A按钮）
	IsFrame    bool      `db:"is_frame"`    // 是否为外链（0是 1否）
	IsVisible  bool      `db:"is_visible"`  // 菜单状态（0显示 1隐藏）
	Icon       string    `db:"icon"`        // 菜单图标
	ReqMethod  string    `db:"req_method"`  // 请求方法
	Status     string    `db:"status"`      // 状态（默认0）
	DelFlag    string    `db:"del_flag"`    // 删除标记
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateBy   string    `db:"create_by"`   // 创建人员
	UpdateBy   string    `db:"update_by"`   // 更新人员
	Remark     string    `db:"remark"`      // 备注
}

type AgoNotice struct {
	ID            int64     `db:"id"`             // 唯一标识
	NoticeTitle   string    `db:"notice_title"`   // 标题
	NoticeType    string    `db:"notice_type"`    // 公告类型（1通知 2公告）
	NoticeContent string    `db:"notice_content"` // 公告内容
	Status        string    `db:"status"`         // 状态（默认0）
	DelFlag       string    `db:"del_flag"`       // 删除标记
	CreateTime    time.Time `db:"create_time"`    // 创建时间
	UpdateTime    time.Time `db:"update_time"`    // 更新时间
	CreateBy      string    `db:"create_by"`      // 创建人员
	UpdateBy      string    `db:"update_by"`      // 更新人员
	Remark        string    `db:"remark"`         // 备注
}

type AgoOperationLog struct {
	ID              int64     `db:"id"`               // 唯一标识
	BusinessType    string    `db:"business_type"`    // 业务类型
	BusinessTitle   string    `db:"business_title"`   // 业务内容
	InvokeMethod    string    `db:"invoke_method"`    // 方法名称
	RequestMethod   string    `db:"request_method"`   // 请求方式
	RequestUrl      string    `db:"request_url"`      // 请求URL
	ClientType      string    `db:"client_type"`      // 操作类别（0其它 1后台用户 2手机端用户）
	ClientIp        string    `db:"client_ip"`        // 主机地址
	ClientLocation  string    `db:"client_location"`  // 操作地点
	ClientParam     string    `db:"client_param"`     // 请求参数
	OperationType   string    `db:"operation_type"`   // 操作类型（0其它 1新增 2修改 3删除）
	OperationResult string    `db:"operation_result"` // 返回参数
	ErrorMsg        string    `db:"error_msg"`        // 错误消息
	Status          string    `db:"status"`           // 状态（0正常 1异常）
	CreateBy        string    `db:"create_by"`        // 操作人员
	CreateTime      time.Time `db:"create_time"`      // 创建时间
	Remark          string    `db:"remark"`           // 备注
	DeptName        string    `db:"dept_name"`        // 部门名称
}

type AgoPost struct {
	ID         int64     `db:"id"`          // 唯一标识
	PostName   string    `db:"post_name"`   // 岗位名称
	OrderNum   int32     `db:"order_num"`   // 排序编号
	Status     string    `db:"status"`      // 状态（默认0）
	DelFlag    string    `db:"del_flag"`    // 删除标记
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateBy   string    `db:"create_by"`   // 创建人员
	UpdateBy   string    `db:"update_by"`   // 更新人员
	Remark     string    `db:"remark"`      // 备注
}

type AgoRole struct {
	ID         int64     `db:"id"`          // 唯一标识
	RoleName   string    `db:"role_name"`   // 角色名称
	RoleKey    string    `db:"role_key"`    // 角色标识
	OrderNum   int32     `db:"order_num"`   // 排序编号
	DataScope  string    `db:"data_scope"`  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     string    `db:"status"`      // 状态（默认0）
	DelFlag    string    `db:"del_flag"`    // 删除标记
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateBy   string    `db:"create_by"`   // 创建人员
	UpdateBy   string    `db:"update_by"`   // 更新人员
	Remark     string    `db:"remark"`      // 备注
}

type AgoRoleDept struct {
	RoleID int64 `db:"role_id"` // 角色
	DeptID int64 `db:"dept_id"` // 部门标识
}

type AgoRoleMenu struct {
	RoleID int64 `db:"role_id"` // 角色标识
	MenuID int64 `db:"menu_id"` // 菜单标识
}

type AgoSession struct {
	ID           uuid.UUID `db:"id"`            // 唯一标识
	UserName     string    `db:"user_name"`     // 用户名
	RealName     string    `db:"real_name"`     // 真实姓名
	RefreshToken string    `db:"refresh_token"` // 刷新秘钥
	UserAgent    string    `db:"user_agent"`    // 请求信息
	ClientIp     string    `db:"client_ip"`     // 请求地址
	IsBlocked    bool      `db:"is_blocked"`    // 是否阻止
	ExpiresAt    time.Time `db:"expires_at"`    // 过期时间
	CreateAt     time.Time `db:"create_at"`     // 创建时间
	Remark       string    `db:"remark"`        // 备注
}

type AgoUser struct {
	ID         int64     `db:"id"`          // 唯一标识
	DeptID     int64     `db:"dept_id"`     // 部门编号
	UserName   string    `db:"user_name"`   // 登录名称
	RealName   string    `db:"real_name"`   // 真实姓名
	Mobile     string    `db:"mobile"`      // 手机号码
	Email      string    `db:"email"`       // 用户邮箱
	Password   string    `db:"password"`    // 用户密码
	Sex        string    `db:"sex"`         // 用户性别（0男 1女 2未知）
	Avatar     string    `db:"avatar"`      // 用户头像
	Posts      string    `db:"posts"`       // 岗位编号数组
	Status     string    `db:"status"`      // 状态（默认0）
	DelFlag    string    `db:"del_flag"`    // 删除标记
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
	CreateBy   string    `db:"create_by"`   // 创建人员
	UpdateBy   string    `db:"update_by"`   // 更新人员
	Remark     string    `db:"remark"`      // 备注
}

type AgoUserRole struct {
	UserID int64 `db:"user_id"` // 用户标识
	RoleID int64 `db:"role_id"` // 角色标识
}
