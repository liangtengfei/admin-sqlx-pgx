package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"time"
)

func sessionCreateSQL(req request.SessionCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert(TBNameRole).
		Columns("user_name", "real_name", "user_agent", "client_ip", "refresh_token", "is_blocked", "expires_at", "create_at", "remark").
		Values(req.UserName, req.RealName, req.UserAgent, req.ClientIp, req.RefreshToken, req.IsBlocked, req.ExpiresAt, time.Now(), req.Remark).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) SessionCreate(ctx context.Context, req request.SessionCreateRequest, username string) (int64, error) {
	sql, args, err := sessionCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func sessionUpdateSQL(req request.SessionUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update(TBNameRole).
		Set("is_blocked", req.IsBlocked).
		Set("remark", req.Remark)
	//sql = sql.Set("update_by", username)
	//sql = sql.Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.ID}).ToSql()
}

func (store *SQLStore) SessionUpdate(ctx context.Context, req request.SessionUpdateRequest, username string) (int64, error) {
	sql, args, err := sessionUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) SessionDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder(TBNameSession, id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) SessionDetail(ctx context.Context, id uuid.UUID) (AgoSession, error) {
	var result AgoSession

	sql, args, err := DetailUUIDSQLBuilder(TBNameSession, id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func sessionPageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder(TBNameSession).Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		sql = sql.Where(sq.Or{
			sq.Like{"session_name": fmt.Sprint("%", req.Keyword, "%")},
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

func (store *SQLStore) SessionPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoSession, error) {
	var result []AgoSession
	var total int64

	fail := func(err error) (int64, []AgoSession, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := sessionPageAndKeywordSQL(req)
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

func (store *SQLStore) SessionList(ctx context.Context) ([]AgoSession, error) {
	var result []AgoSession

	sql, args, err := SQLBuilder().Select("*").From(TBNameSession).Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) SessionListByIds(ctx context.Context, ids string) ([]AgoSession, error) {
	var result []AgoSession

	sql, args, err := SQLBuilder().Select("*").From(TBNameSession).
		Where("id = ANY(STRING_TO_ARRAY(?, ',')::int8[])", ids).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}
