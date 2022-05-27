package request

import "time"

type OperationLogCreateRequest struct {
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
