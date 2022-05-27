package response

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// OperationLogResponse 操作日志记录
type OperationLogResponse struct {
	ID              int64  `json:"id"`              // 日志主键
	BusinessType    string `json:"businessType"`    // 业务类型
	BusinessTitle   string `json:"businessTitle"`   // 业务内容
	OperationType   int32  `json:"operationType"`   // 操作类型（0其它 1新增 2修改 3删除）
	Method          string `json:"method"`          // 方法名称
	RequestMethod   string `json:"requestMethod"`   // 请求方式
	RequestUrl      string `json:"requestUrl"`      // 请求URL
	ClientType      int32  `json:"clientType"`      // 操作类别（0其它 1后台用户 2手机端用户）
	ClientIp        string `json:"clientIp"`        // 主机地址
	ClientLocation  string `json:"clientLocation"`  // 操作地点
	ClientParam     string `json:"clientParam"`     // 请求参数
	OperationResult string `json:"operationResult"` // 返回参数
	ErrorMsg        string `json:"errorMsg"`        // 错误消息
	Status          string `json:"status"`          // 状态（0正常 1异常）
	DeptName        string `json:"deptName"`        // 部门名称
	CreateBy        string `json:"createBy"`        // 创建人员
	CreateTimeStr   string `json:"createTime"`      // 创建时间
	Remark          string `json:"remark"`          // 备注
}

func (res *OperationLogResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}
