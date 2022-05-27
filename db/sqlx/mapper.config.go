package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"time"
)

func configCreateSQL(req request.SysConfigCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameConfig).
		Columns("config_name", "config_key", "config_value", "create_time", "create_by", "remark").
		Values(req.ConfigName, req.ConfigKey, req.ConfigValue, time.Now(), username, req.Remark).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) ConfigCreate(ctx context.Context, req request.SysConfigCreateRequest, username string) (int64, error) {
	sql, args, err := configCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func configCreateSQLBatch(reqs []request.SysConfigCreateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Insert(TBNameConfig).
		Columns("config_name", "config_key", "config_value", "create_time", "create_by", "remark")
	for _, req := range reqs {
		sql = sql.Values(req.ConfigName, req.ConfigKey, req.ConfigValue, time.Now(), username, req.Remark)
	}

	return sql.ToSql()
}

func (store *SQLStore) ConfigCreateBatch(ctx context.Context, req []request.SysConfigCreateRequest, username string) (int64, error) {
	sql, args, err := configCreateSQLBatch(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func configUpdateSQL(req request.SysConfigUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameConfig).
		Set("config_name", req.ConfigName).
		Set("config_value", req.ConfigValue).
		Set("status", req.Status).
		Set("remark", req.Remark)
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) ConfigUpdate(ctx context.Context, req request.SysConfigUpdateRequest, username string) (int64, error) {
	sql, args, err := configUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) ConfigDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNameConfig, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) ConfigDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder(TBNameConfig, id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) ConfigDetail(ctx context.Context, id int64) (AgoConfig, error) {
	var result AgoConfig

	sql, args, err := DetailSQLBuilder(TBNameConfig, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func configPageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameConfig).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"config_name": fmt.Sprint("%", req.Keyword, "%")},
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

func (store *SQLStore) ConfigPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoConfig, error) {
	var result []AgoConfig
	var total int64

	fail := func(err error) (int64, []AgoConfig, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := configPageAndKeywordSQL(req)
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

func (store *SQLStore) ConfigList(ctx context.Context) ([]AgoConfig, error) {
	var result []AgoConfig

	sql, args, err := SQLBuilder().Select("*").From(TBNameConfig).Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) ConfigListByIds(ctx context.Context, ids string) ([]AgoConfig, error) {
	var result []AgoConfig

	sql, args, err := SQLBuilder().Select("*").From(TBNameConfig).
		Where("id = ANY(STRING_TO_ARRAY(?, ',')::int8[])", ids).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}

func (store *SQLStore) ConfigCountByKey(ctx context.Context, key string) (int64, error) {
	sql, args, err := SQLBuilder().Select("count(*)").From(TBNameConfig).Where(sq.Eq{"ConfigKey": key}).ToSql()
	if err != nil {
		return 0, err
	}
	var total int64
	err = store.db.GetContext(ctx, &total, sql, args...)
	return total, err
}
