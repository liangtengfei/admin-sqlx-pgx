package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/global/consts"
	"time"
)

func menuCreateSQL(req request.MenuCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameMenu).
		Columns("menu_name", "menu_key", "parent_id", "path", "menu_type", "is_frame", "is_visible", "icon", "req_method", "create_time", "create_by", "remark").
		Values(req.MenuName, req.MenuKey, req.ParentID, req.Path, req.MenuType, req.IsFrame, req.Visible, req.Icon, req.RequestMethod, time.Now(), username, req.Remark).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) MenuCreate(ctx context.Context, req request.MenuCreateRequest, username string) (int64, error) {
	sql, args, err := menuCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func menuUpdateSQL(req request.MenuUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameMenu).
		Set("menu_name", req.MenuName).
		Set("menu_key", req.MenuKey).
		Set("parent_id", req.ParentID).
		Set("path", req.Path).
		Set("menu_type", req.MenuType).
		Set("is_frame", req.IsFrame).
		Set("is_visible", req.Visible).
		Set("icon", req.Icon).
		Set("req_method", req.RequestMethod).
		Set("status", req.Status).
		Set("remark", req.Remark)
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) MenuUpdate(ctx context.Context, req request.MenuUpdateRequest, username string) (int64, error) {
	sql, args, err := menuUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) MenuDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNameMenu, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) MenuDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder(TBNameMenu, id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) MenuDetail(ctx context.Context, id int64) (AgoMenu, error) {
	var result AgoMenu

	sql, args, err := DetailSQLBuilder(TBNameMenu, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func menuPageAndKeywordSQL(ctx context.Context, req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameMenu).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"menu_name": fmt.Sprint("%", req.Keyword, "%")},
			sq.Like{"menu_key": fmt.Sprint("%", req.Keyword, "%")},
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

func (store *SQLStore) MenuPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoMenu, error) {
	var result []AgoMenu
	var total int64

	fail := func(err error) (int64, []AgoMenu, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := menuPageAndKeywordSQL(ctx, req)
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

func (store *SQLStore) MenuList(ctx context.Context) ([]AgoMenu, error) {
	var result []AgoMenu

	sql, args, err := SQLBuilder().Select("*").From(TBNameMenu).Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) MenuListByRoleId(ctx context.Context, id int64) ([]AgoMenu, error) {
	var result []AgoMenu

	sql, args, err := SQLBuilder().Select("*").From(TBNameMenu).
		Where("id IN (SELECT menu_id FROM ago_role_menu WHERE role_id = ?)", id).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}
