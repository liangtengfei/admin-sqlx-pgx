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

func postCreateSQL(req request.PostCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNamePost).
		Columns("post_name", "order_num", "create_time", "create_by", "remark").
		Values(req.PostName, req.OrderNum, time.Now(), username, req.Remark).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) PostCreate(ctx context.Context, req request.PostCreateRequest, username string) (int64, error) {
	sql, args, err := postCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func postUpdateSQL(req request.PostUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNamePost).
		Set("post_name", req.PostName).
		Set("order_num", req.OrderNum).
		Set("status", req.Status).
		Set("remark", req.Remark)
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) PostUpdate(ctx context.Context, req request.PostUpdateRequest, username string) (int64, error) {
	sql, args, err := postUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) PostDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNamePost, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) PostDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder(TBNamePost, id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) PostDetail(ctx context.Context, id int64) (internal.AgoPost, error) {
	var result internal.AgoPost

	sql, args, err := DetailSQLBuilder(TBNamePost, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func postPageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNamePost).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"post_name": fmt.Sprint("%", req.Keyword, "%")},
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

func (store *SQLStore) PostPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoPost, error) {
	var result []internal.AgoPost
	var total int64

	fail := func(err error) (int64, []internal.AgoPost, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := postPageAndKeywordSQL(req)
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

func (store *SQLStore) PostList(ctx context.Context) ([]internal.AgoPost, error) {
	var result []internal.AgoPost

	sql, args, err := SQLBuilder().Select("*").From(TBNamePost).Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) PostListByIds(ctx context.Context, ids string) ([]internal.AgoPost, error) {
	var result []internal.AgoPost

	sql, args, err := SQLBuilder().Select("*").From(TBNamePost).
		Where("id = ANY(STRING_TO_ARRAY(?, ',')::int8[])", ids).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}
