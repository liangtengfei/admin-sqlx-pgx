package service

import (
	"errors"
	"github.com/jmoiron/sqlx"
	db "study.com/demo-sqlx-pgx/db/sqlx"
)

var (
	store db.Store
)

const (
	BizTitleUser         = "系统用户"
	BizTitleRole         = "系统角色"
	BizTitleDept         = "系统部门"
	BizTitleMenu         = "系统菜单"
	BizTitlePost         = "系统岗位"
	BizTitleConfig       = "系统配置"
	BizTitleNotice       = "公告提醒"
	BizTitleDictType     = "字典类型"
	BizTitleDictData     = "字典数据"
	BizTitleSession      = "登录会话"
	BizTitleOperationLog = "系统日志"
)

const (
	OperationTypeQuery  = "查询"
	OperationTypeCreate = "新增"
	OperationTypeUpdate = "更新"
	OperationTypeDelete = "删除"
	OperationTypePage   = "分页查询"
	OperationTypeList   = "列表查询"
	OperationTypeDetail = "详情"
)

// 对用户隐藏数据库的相关错误
var (
	ErrQuery  = errors.New("查询失败")
	ErrPage   = errors.New("分页查询失败")
	ErrCreate = errors.New("新增失败")
	ErrUpdate = errors.New("更新失败")
	ErrDelete = errors.New("删除失败")
	ErrNoRows = errors.New("未查询到信息")
)

func InitService(conn *sqlx.DB) {
	store = db.NewStore(conn)
}

func attachPageParams(pageNum, pageSize int32) (int32, int32) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize > 10000 {
		pageSize = 10000
	}

	offset := (pageNum - 1) * pageSize
	return offset, pageSize
}
