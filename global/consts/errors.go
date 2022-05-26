package consts

import "errors"

// CURD 常用业务状态码
const (
	ValidatorParamsCheckFailMsg string = "参数校验失败"

	ServerOccurredErrorMsg string = "服务器内部发生代码执行错误, "

	JwtTokenFormatErrMsg string = "提交的 token 格式错误"            //提交的 token 格式错误
	JwtTokenMustValid    string = "token为必填项,请在请求header部分提交!" //提交的 token 格式错误

	CurdCreatFailMsg        string = "新增失败"
	CurdUpdateFailMsg       string = "更新失败"
	CurdDeleteFailMsg       string = "删除失败"
	CurdSelectFailMsg       string = "未查询到数据"
	CurdRegisterFailMsg     string = "注册失败"
	CurdLoginFailMsg        string = "登录失败"
	CurdRefreshTokenFailMsg string = "刷新Token失败"

	FilesUploadFailMsg            string = "文件上传失败, 获取上传文件发生错误!"
	FilesUploadMoreThanMaxSizeMsg string = "长传文件超过系统设定的最大值,系统允许的最大值（M）："
	FilesUploadMimeTypeFailMsg    string = "文件mime类型不允许"

	CaptchaGetParamsInvalidMsg   string = "获取验证码：提交的验证码参数无效,请检查验证码ID以及文件名后缀是否完整"
	CaptchaCheckParamsInvalidMsg string = "校验验证码：提交的参数无效，请检查 【验证码ID、验证码值】 提交时的键名是否与配置项一致"
	CaptchaCheckOkMsg            string = "验证码校验通过"
	CaptchaCheckFailMsg          string = "验证码校验失败"
)

var (
	ErrorNotFind = errors.New("未查询到数据")
)
