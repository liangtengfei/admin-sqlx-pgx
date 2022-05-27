package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"time"
)

func roleCreateSQL(req request.RoleCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameRole).
		Columns("role_name", "role_key", "order_num", "data_scope", "remark", "create_time", "create_by").
		Values(req.RoleName, req.RoleKey, req.OrderNum, req.DataScope, req.Remark, time.Now(), username).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) RoleCreate(ctx context.Context, req request.RoleCreateRequest, username string) (int64, error) {
	sql, args, err := roleCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	// 事务执行
	tx, err := store.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	// 执行操作
	var id int64
	err = tx.QueryRowx(sql, args...).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if len(req.DeptIds) > 0 {
		err = roleDeptRelate(tx, id, req.DeptIds)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}

	if len(req.MenuIds) > 0 {
		err = roleMenuRelate(tx, id, req.MenuIds)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}

	return id, nil
}

func roleUpdateSQL(req request.RoleUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameRole).
		Set("role_name", req.RoleName).
		Set("order_num", req.OrderNum).
		Set("data_scope", req.DataScope).
		Set("status", req.Status).
		Set("remark", req.Remark)
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) RoleUpdate(ctx context.Context, req request.RoleUpdateRequest, username string) (int64, error) {
	var result int64

	sql, args, err := roleUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	// 事务执行
	tx, err := store.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	// 执行操作
	res, err := tx.Exec(sql, args...)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	result, err = res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if len(req.DeptIds) > 0 {
		err = roleDeptRelate(tx, req.ID, req.DeptIds)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}

	if len(req.MenuIds) > 0 {
		err = roleMenuRelate(tx, req.ID, req.MenuIds)
		if err != nil {
			_ = tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	return result, err
}

func (store *SQLStore) RoleDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNameRole, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) RoleDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder(TBNameRole, id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) RoleDetail(ctx context.Context, id int64) (AgoRole, error) {
	var result AgoRole

	sql, args, err := DetailSQLBuilder(TBNameRole, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) RoleListByUserId(id int64) ([]AgoRole, error) {
	var result []AgoRole

	sql, args, err := SQLBuilder().Select("*").From(TBNameRole).
		Where("id IN (SELECT role_id FROM ago_user_role WHERE user_id = ?)", id).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.Select(&result, sql, args...)
	return result, err
}

func roleDeptRelate(tx *sqlx.Tx, id int64, relateIds []int64) error {
	return relateDataReset(tx, id, relateIds, "role_id", "dept_id", TBNameRoleDept)
}

func roleMenuRelate(tx *sqlx.Tx, id int64, relateIds []int64) error {
	return relateDataReset(tx, id, relateIds, "role_id", "menu_id", TBNameRoleMenu)
}

func rolePageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameRole).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"role_name": fmt.Sprint("%", req.Keyword, "%")},
			sq.Like{"role_key": fmt.Sprint("%", req.Keyword, "%")},
		})
	}

	// 此处截取COUNT SQL
	countSQL, _, err = sql.ToSql()
	if err != nil {
		return
	}
	countSQL = SQLCount(countSQL)

	//分页
	sql = sql.Offset(req.GetOffset()).Limit(req.GetLimit())

	//排序
	if req.SortField != "" && req.SortOrder != "" {
		sql = sql.OrderBy(req.SortField + " " + req.SortOrder)
	} else {
		sql = sql.OrderBy("create_time DESC")
	}

	querySQL, args, err = sql.ToSql()

	return querySQL, countSQL, args, err
}

func (store *SQLStore) RolePage(ctx context.Context, req request.PaginationRequest) (int64, []AgoRole, error) {
	var result []AgoRole
	var total int64

	fail := func(err error) (int64, []AgoRole, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := rolePageAndKeywordSQL(req)
	if err != nil {
		return fail(err)
	}

	err = store.db.GetContext(ctx, &total, countSQL, args...)
	if err != nil {
		return fail(err)
	}

	err = store.db.SelectContext(ctx, &result, pageSQL, args...)
	if err != nil {
		return fail(err)
	}

	return total, result, nil
}

func (store *SQLStore) RoleCountByKey(ctx context.Context, roleKey string) (int64, error) {
	sql, args, err := SQLBuilder().Select("count(*)").From(TBNameRole).Where(sq.Eq{"role_key": roleKey}).ToSql()
	if err != nil {
		return 0, err
	}
	var total int64
	err = store.db.GetContext(ctx, &total, sql, args...)
	return total, err
}
