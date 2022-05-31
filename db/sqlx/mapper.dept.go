package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/db/sqlx/internal"
	"study.com/demo-sqlx-pgx/global/consts"
	"time"
)

func deptCreateSQL(req request.DeptCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameDept).
		Columns("parent_id", "ancestors", "dept_name", "dept_code", "order_num", "create_time", "create_by", "remark").
		Values(req.ParentID, req.Ancestors, req.DeptName, req.DeptCode, req.OrderNum, time.Now(), username, req.Remark).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) DeptCreate(ctx context.Context, req request.DeptCreateRequest, username string) (int64, error) {
	sql, args, err := deptCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func deptUpdateSQL(req request.DeptUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameDept).
		Set("parent_id", req.ParentID).
		Set("dept_name", req.DeptName).
		Set("dept_code", req.DeptCode).
		Set("order_num", req.OrderNum).
		Set("ancestors", req.Ancestors).
		Set("remark", req.Remark)
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) DeptUpdate(ctx context.Context, req request.DeptUpdateRequest, username string) (int64, error) {
	sql, args, err := deptUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) DeptDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNameDept, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) DeptDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder(TBNameDept, id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) DeptDetail(ctx context.Context, id int64) (internal.AgoDept, error) {
	var result internal.AgoDept

	sql, args, err := DetailSQLBuilder(TBNameDept, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func deptPageAndKeywordSQL(ctx context.Context, req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameDept).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"dept_name": fmt.Sprint("%", req.Keyword, "%")},
			sq.Like{"dept_code": fmt.Sprint("%", req.Keyword, "%")},
		})
	}

	scope := ctx.Value(consts.ScopeDataKey).(request.DataScopeRequest)
	sql = DataScopeSQLBuilder(sql, scope.Scope, scope.Params)

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

func (store *SQLStore) DeptPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoDept, error) {
	var result []internal.AgoDept
	var total int64

	fail := func(err error) (int64, []internal.AgoDept, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := deptPageAndKeywordSQL(ctx, req)
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

func (store *SQLStore) DeptList(ctx context.Context) ([]internal.AgoDept, error) {
	var result []internal.AgoDept
	sql := SQLBuilder().Select("*").From(TBNameDept).Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	})

	scope := ctx.Value(consts.ScopeDataKey).(request.DataScopeRequest)
	sql = DataScopeSQLBuilder(sql, scope.Scope, scope.Params)

	querySQL, args, err := sql.ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, querySQL, args...)

	return result, err
}

func (store *SQLStore) DeptListByRoleId(ctx context.Context, id int64) ([]internal.AgoDept, error) {
	var result []internal.AgoDept

	sql, args, err := SQLBuilder().Select("*").From(TBNameDept).
		Where("id IN (SELECT dept_id FROM ago_role_dept WHERE role_id = ?)", id).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}
