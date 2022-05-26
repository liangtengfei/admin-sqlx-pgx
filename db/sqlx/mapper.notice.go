package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"time"
)

func noticeCreateSQL(req request.NoticeCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameRole).
		Columns("notice_title", "notice_type", "notice_content", "create_time", "create_by", "remark").
		Values(req.NoticeTitle, req.NoticeType, req.NoticeContent, time.Now(), username, req.Remark).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) NoticeCreate(ctx context.Context, req request.NoticeCreateRequest, username string) (int64, error) {
	sql, args, err := noticeCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func noticeUpdateSQL(req request.NoticeUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameRole).
		Set("notice_title", req.NoticeTitle).
		Set("notice_type", req.NoticeType).
		Set("notice_content", req.NoticeContent).
		Set("status", req.Status).
		Set("remark", req.Remark)
	sql = sql.Set("update_by", username)
	sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) NoticeUpdate(ctx context.Context, req request.NoticeUpdateRequest, username string) (int64, error) {
	sql, args, err := noticeUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) NoticeDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNameNotice, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) NoticeDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder(TBNameNotice, id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) NoticeDetail(ctx context.Context, id int64) (AgoNotice, error) {
	var result AgoNotice

	sql, args, err := DetailSQLBuilder(TBNameNotice, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func noticePageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameNotice).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"notice_name": fmt.Sprint("%", req.Keyword, "%")},
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

func (store *SQLStore) NoticePage(ctx context.Context, req request.PaginationRequest) (int64, []AgoNotice, error) {
	var result []AgoNotice
	var total int64

	fail := func(err error) (int64, []AgoNotice, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := noticePageAndKeywordSQL(req)
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

func (store *SQLStore) NoticeList(ctx context.Context) ([]AgoNotice, error) {
	var result []AgoNotice

	sql, args, err := SQLBuilder().Select("*").From(TBNameNotice).Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) NoticeListByIds(ctx context.Context, ids string) ([]AgoNotice, error) {
	var result []AgoNotice

	sql, args, err := SQLBuilder().Select("*").From(TBNameNotice).
		Where("id = ANY(STRING_TO_ARRAY(?, ',')::int8[])", ids).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}
