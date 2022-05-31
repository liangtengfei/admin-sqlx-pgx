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

func dictDataCreateSQL(req request.DictDataCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameDictData).
		Columns("dict_type", "dict_label", "dict_value", "order_num", "css_class", "list_class", "create_time", "create_by", "remark").
		Values(req.DictType, req.DictLabel, req.DictValue, req.OrderNum, req.CssClass, req.ListClass, time.Now(), username, req.Remark).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) DictDataCreate(ctx context.Context, req request.DictDataCreateRequest, username string) (int64, error) {
	sql, args, err := dictDataCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func dictDataUpdateSQL(req request.DictDataUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameDictData).
		Set("dict_name", req.DictLabel).
		Set("dict_name", req.DictValue).
		Set("dict_name", req.OrderNum).
		Set("dict_name", req.CssClass).
		Set("dict_type", req.ListClass).
		Set("status", req.Status).
		Set("remark", req.Remark)
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) DictDataUpdate(ctx context.Context, req request.DictDataUpdateRequest, username string) (int64, error) {
	sql, args, err := dictDataUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) DictDataDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNameDictData, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) DictDataDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder(TBNameDictData, id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) DictDataDetail(ctx context.Context, id int64) (internal.AgoDictData, error) {
	var result internal.AgoDictData
	sql, args, err := DetailSQLBuilder(TBNameDictData, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func dictDataPageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameDictData).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"dictData_name": fmt.Sprint("%", req.Keyword, "%")},
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

func (store *SQLStore) DictDataPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoDictData, error) {
	var result []internal.AgoDictData
	var total int64

	fail := func(err error) (int64, []internal.AgoDictData, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := dictDataPageAndKeywordSQL(req)
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

func (store *SQLStore) DictDataList(ctx context.Context) ([]internal.AgoDictData, error) {
	var result []internal.AgoDictData

	sql, args, err := SQLBuilder().Select("*").From(TBNameDictData).Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) DictDataListByIds(ctx context.Context, ids string) ([]internal.AgoDictData, error) {
	var result []internal.AgoDictData

	sql, args, err := SQLBuilder().Select("*").From(TBNameDictData).
		Where("id = ANY(STRING_TO_ARRAY(?, ',')::int8[])", ids).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}
