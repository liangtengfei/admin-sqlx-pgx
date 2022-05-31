package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/db/sqlx/internal"
	"time"
)

func dictTypeCreateSQL(req request.DictTypeCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameDictType).
		Columns("dict_name", "dict_type", "create_time", "create_by", "remark").
		Values(req.DictName, req.DictType, time.Now(), username, req.Remark).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) DictTypeCreate(ctx context.Context, req request.DictTypeCreateRequest, username string) (int64, error) {
	sql, args, err := dictTypeCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func dictTypeUpdateSQL(req request.DictTypeUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameDictType).
		Set("dict_name", req.DictName).
		Set("dict_type", req.DictType).
		Set("status", req.Status).
		Set("remark", req.Remark)
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) DictTypeUpdate(ctx context.Context, req request.DictTypeUpdateRequest, username string) (int64, error) {
	sql, args, err := dictTypeUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) DictTypeDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNameDictType, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) DictTypeDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder(TBNameDictType, id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) DictTypeDetail(ctx context.Context, id int64) (internal.AgoDictType, error) {
	var result internal.AgoDictType

	sql, args, err := DetailSQLBuilder(TBNameDictType, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func dictTypePageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameDictType).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"dictType_name": fmt.Sprint("%", req.Keyword, "%")},
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

func (store *SQLStore) DictTypePage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoDictType, error) {
	var result []internal.AgoDictType
	var total int64

	fail := func(err error) (int64, []internal.AgoDictType, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := dictTypePageAndKeywordSQL(req)
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

func (store *SQLStore) DictTypeList(ctx context.Context) ([]internal.AgoDictType, error) {
	var result []internal.AgoDictType

	sql, args, err := SQLBuilder().Select("*").From(TBNameDictType).Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) DictTypeListByIds(ctx context.Context, ids string) ([]internal.AgoDictType, error) {
	var result []internal.AgoDictType

	sql, args, err := SQLBuilder().Select("*").From(TBNameDictType).
		Where("id = ANY(STRING_TO_ARRAY(?, ',')::int8[])", ids).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}
