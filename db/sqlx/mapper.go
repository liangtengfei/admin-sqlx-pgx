package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/global/consts"
	"time"
)

const (
	TBNameUser         = "ago_user"
	TBNameDept         = "ago_dept"
	TBNameRole         = "ago_role"
	TBNameMenu         = "ago_menu"
	TBNamePost         = "ago_post"
	TBNameNotice       = "ago_notice"
	TBNameConfig       = "ago_config"
	TBNameDictType     = "ago_dict_type"
	TBNameDictData     = "ago_dict_data"
	TBNameOperationLog = "ago_operation_log"
	TBNameSession      = "ago_session"
	TBNameRoleMenu     = "ago_role_menu"
	TBNameRoleDept     = "ago_role_dept"
	TBNameUserRole     = "ago_user_role"
)

func SQLBuilder(format ...string) sq.StatementBuilderType {
	var dbType string
	if len(format) > 0 {
		dbType = format[0]
	} else {
		dbType = global.Config.Server.DBDriver
	}

	if dbType == "mysql" {
		return sq.StatementBuilder.PlaceholderFormat(sq.Question)
	} else if dbType == "postgresql" {
		return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	} else if dbType == "pgx" {
		return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	}
	return sq.StatementBuilder.PlaceholderFormat(sq.Question)
}

func baseQuerySQLBuilder(name string) sq.SelectBuilder {
	return SQLBuilder().Select("*").From(name)
}

func SQLCount(sql string) string {
	if strings.Index(sql, "*") > 0 {
		return strings.Replace(sql, "*", "count(*)", 1)
	}
	return sql
}

func DetailSQLBuilder(tableName string, id int64) (querySQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(tableName).
		Where(sq.Eq{"id": id}).
		Where(sq.Eq{"status": "0"}).
		Where(sq.Eq{"del_flag": "N"}).
		Limit(1)
	querySQL, args, err = sql.ToSql()
	return
}

func DetailUUIDSQLBuilder(tableName string, id uuid.UUID) (querySQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(tableName).
		Where(sq.Eq{"id": id}).
		Where(sq.Eq{"status": "0"}).
		Where(sq.Eq{"del_flag": "N"}).
		Limit(1)
	querySQL, args, err = sql.ToSql()
	return
}

func DeleteSQLBuilder(tableName string, id int64) (querySQL string, args []interface{}, err error) {
	sql := SQLBuilder().Delete(tableName).Where(sq.Eq{"id": id})
	querySQL, args, err = sql.ToSql()
	return
}

func DeleteFakeSQLBuilder(tableName string, id int64, username string) (querySQL string, args []interface{}, err error) {
	sql := SQLBuilder().
		Update(tableName).
		Set("del_flag", "Y").
		Set("update_by", username).
		Set("update_time", time.Now()).
		Where(sq.Eq{"id": id})
	querySQL, args, err = sql.ToSql()
	return
}

// 重置关联数据 先删除 再新增
func relateDataReset(tx *sqlx.Tx, id int64, relateIds []int64, mainField, relateField string, tableName string) error {
	// 删除关联
	delSQL, delArgs, err := SQLBuilder().Delete(tableName).Where(sq.Eq{mainField: id}).ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(delSQL, delArgs...)
	if err != nil {
		return err
	}

	// 填充参数
	build := SQLBuilder().Insert(tableName).Columns(mainField, relateField)
	for _, relateId := range relateIds {
		build = build.Values(id, relateId)
	}

	// 新增关联
	insertSQL, insertArgs, err := build.ToSql()

	if err != nil {
		return err
	}
	_, err = tx.Exec(insertSQL, insertArgs...)
	return err
}

// DataScopeSQLBuilder 数据过滤构建
func DataScopeSQLBuilder(sql sq.SelectBuilder, scope string, data []interface{}) sq.SelectBuilder {
	if scope == consts.ScopeDataAll.String() {
		return sql
	} else if scope == consts.ScopeCustom.String() {
		return sql.Where("id IN (SELECT dept_id FROM ago_role_dept WHERE role_id = ?)", data[0])
	} else if scope == consts.ScopeDept.String() {
		return sql.Where("id = ?", data[0])
	} else if scope == consts.ScopeDeptChild.String() {
		return sql.Where("id = ? OR ? = ANY(STRING_TO_ARRAY(ancestors, ',')::int8[])", data[0], data[0])
	}
	return sql.Where("create_by = ?", data[0])
}
