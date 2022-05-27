package response

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"study.com/demo-sqlx-pgx/utils/valid"
)

const (
	SUCCESS = http.StatusOK
	ERROR   = http.StatusInternalServerError
)

type RestRes struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"message" example:"status ok"`
	Data interface{} `json:"data,omitempty"`
}

type PaginationRes struct {
	Total int64
	RestRes
}

func response(code int, msg string, data interface{}) RestRes {
	return RestRes{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Success(ctx *gin.Context) {
	ctx.JSON(SUCCESS, response(SUCCESS, "操作成功", gin.H{}))
}
func SuccessData(ctx *gin.Context, data interface{}) {
	ctx.JSON(SUCCESS, response(SUCCESS, "操作成功", data))
}

func SuccessPage(ctx *gin.Context, total int64, data interface{}) {
	res := PaginationRes{
		Total: total,
		RestRes: RestRes{
			Code: SUCCESS,
			Msg:  "操作成功",
			Data: data,
		},
	}

	ctx.JSON(SUCCESS, res)
}

func Error(ctx *gin.Context) {
	ctx.JSON(ERROR, response(ERROR, "操作失败", gin.H{}))
	ctx.Abort()
}

// ErrorValid 请求参数验证错误时
func ErrorValid(ctx *gin.Context, err error) {
	var buff bytes.Buffer

	//避免格式错误
	switch err.(type) {
	case validator.ValidationErrors:
		m := valid.Translate(err)
		msgs := make([]string, 0)
		for _, v := range m {
			msgs = append(msgs, v...)
		}
		buff.WriteString(strings.Join(msgs, ","))

		ctx.AbortWithStatusJSON(ERROR, response(ERROR, buff.String(), nil))
	default:
		ctx.AbortWithStatusJSON(ERROR, response(ERROR, "参数格式不正确", nil))
	}
}

func ErrorMsg(ctx *gin.Context, msg string) {
	ctx.JSON(ERROR, response(ERROR, msg, gin.H{}))
	ctx.Abort()
}

func ErrorScope(ctx *gin.Context) {
	ctx.JSON(ERROR, response(ERROR, "无数据权限", gin.H{}))
	ctx.Abort()
}

func ErrorCodeMsg(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, response(code, msg, gin.H{}))
	ctx.Abort()
}
