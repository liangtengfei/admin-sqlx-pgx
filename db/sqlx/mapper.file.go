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

func fileCreateSQL(req request.FileCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert("ago_file").
		Columns(
			"file_name",
			"file_path",
			"file_url",
			"file_size",
			"user_id",
			"mime_type",
			"create_time",
			"create_by",
			"remark").
		Values(
			req.FileName,
			req.FilePath,
			req.FileUrl,
			req.FileSize,
			req.UserId,
			req.MimeType,
			time.Now(),
			username,
			req.Remark,
		).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) FileCreate(ctx context.Context, req request.FileCreateRequest, username string) (int64, error) {
	sql, args, err := fileCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func fileCreateSQLBatch(reqs []request.FileCreateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Insert("ago_file").
		Columns(
			"id",
			"file_name",
			"file_path",
			"file_url",
			"file_size",
			"user_id",
			"mime_type",
			"create_time",
			"create_by",
			"remark")
	for _, req := range reqs {
		sql = sql.Values(
			req.Id,
			req.FileName,
			req.FilePath,
			req.FileUrl,
			req.FileSize,
			req.UserId,
			req.MimeType,
			time.Now(),
			username,
			req.Remark,
		)
	}

	return sql.ToSql()
}

func (store *SQLStore) FileCreateBatch(ctx context.Context, req []request.FileCreateRequest, username string) (int64, error) {
	sql, args, err := fileCreateSQLBatch(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func fileUpdateSQL(req request.FileUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update("ago_file").
		Set("file_name", req.FileName).
		Set("file_path", req.FilePath).
		Set("file_url", req.FileUrl).
		Set("file_size", req.FileSize).
		Set("user_id", req.UserId).
		Set("mime_type", req.MimeType).
		Set("remark", req.Remark).
		Set("update_by", username).
		Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.Id}).ToSql()
}

func (store *SQLStore) FileUpdate(ctx context.Context, req request.FileUpdateRequest, username string) (int64, error) {
	sql, args, err := fileUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) FileDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder("ago_file", id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) FileDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder("ago_file", id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) FileDetail(ctx context.Context, id int64) (internal.AgoFile, error) {
	var result internal.AgoFile

	sql, args, err := DetailSQLBuilder("ago_file", id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func filePageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder("ago_file")
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		// 根据实际情况填充
		sql = sql.Where(sq.Or{
			sq.Like{"file_name": fmt.Sprint("%", req.Keyword, "%")},
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

func (store *SQLStore) FilePage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoFile, error) {
	var result []internal.AgoFile
	var total int64

	fail := func(err error) (int64, []internal.AgoFile, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := filePageAndKeywordSQL(req)
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

func (store *SQLStore) FileList(ctx context.Context) ([]internal.AgoFile, error) {
	var result []internal.AgoFile

	sql, args, err := SQLBuilder().Select("*").From("ago_file").Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) FileListByIds(ctx context.Context, ids string) ([]internal.AgoFile, error) {
	var result []internal.AgoFile

	sql, args, err := SQLBuilder().Select("*").From("ago_file").
		Where("id = ANY(STRING_TO_ARRAY(?, ',')::int8[])", ids).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}

// FileCountByKey 根据关键字统计条数（根据情况启用）
//func (store *SQLStore) FileCountByKey(ctx context.Context, key string) (int64, error) {
//	sql, args, err := SQLBuilder().Select("count(*)").From("ago_file").Where(sq.Eq{"FileKey": key}).ToSql()
//	if err != nil {
//		return 0, err
//	}
//	var total int64
//	err = store.db.GetContext(ctx, &total, sql, args...)
//	return total, err
//}
