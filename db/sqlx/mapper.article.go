package db

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"strings"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/db/sqlx/internal"
	"time"
)

func articleCreateSQL(req request.ArticleCreateRequest, username string) (string, []interface{}, error) {
	return SQLBuilder().Insert("cms_article").
		Columns(
			"article_title",
			"article_content",
			"article_tag",
			"summary",
			"create_time",
			"create_by",
			"remark").
		Values(
			req.ArticleTitle,
			req.ArticleContent,
			req.ArticleTag,
			req.Summary,
			time.Now(),
			username,
			req.Remark,
		).
		Suffix("RETURNING \"id\"").
		ToSql()
}

func (store *SQLStore) ArticleCreate(ctx context.Context, req request.ArticleCreateRequest, username string) (int64, error) {
	sql, args, err := articleCreateSQL(req, username)
	if err != nil {
		return 0, err
	}

	var id int64
	err = store.db.QueryRowxContext(ctx, sql, args...).Scan(&id)

	return id, err
}

func articleCreateSQLBatch(reqs []request.ArticleCreateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Insert("cms_article").
		Columns(
			"article_title",
			"article_content",
			"article_tag",
			"summary",
			"create_time",
			"create_by",
			"remark")
	for _, req := range reqs {
		sql = sql.Values(
			req.ArticleTitle,
			req.ArticleContent,
			req.ArticleTag,
			req.Summary,
			time.Now(),
			username,
			req.Remark,
		)
	}

	return sql.ToSql()
}

func (store *SQLStore) ArticleCreateBatch(ctx context.Context, req []request.ArticleCreateRequest, username string) (int64, error) {
	sql, args, err := articleCreateSQLBatch(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func articleUpdateSQL(req request.ArticleUpdateRequest, username string) (string, []interface{}, error) {
	sql := SQLBuilder().Update("cms_article").
		Set("article_title", req.ArticleTitle).
		Set("article_content", req.ArticleContent).
		Set("article_tag", req.ArticleTag).
		Set("summary", req.Summary).
		Set("status", req.Status).
		Set("remark", req.Remark).
		Set("update_by", username).
		Set("update_time", time.Now())
	return sql.Where(sq.Eq{"id": req.Id}).ToSql()
}

func (store *SQLStore) ArticleUpdate(ctx context.Context, req request.ArticleUpdateRequest, username string) (int64, error) {
	sql, args, err := articleUpdateSQL(req, username)
	if err != nil {
		return 0, err
	}

	res, err := store.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (store *SQLStore) ArticleDelete(ctx context.Context, id int64) (int64, error) {
	var result int64

	sql, args, err := DeleteSQLBuilder("cms_article", id)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) ArticleDeleteFake(ctx context.Context, id int64, username string) (int64, error) {
	var result int64

	sql, args, err := DeleteFakeSQLBuilder("cms_article", id, username)
	if err != nil {
		return result, err
	}
	_, err = store.db.ExecContext(ctx, sql, args...)

	return result, err
}

func (store *SQLStore) ArticleDetail(ctx context.Context, id int64) (internal.CmsArticle, error) {
	var result internal.CmsArticle

	sql, args, err := DetailSQLBuilder("cms_article", id)
	if err != nil {
		return result, err
	}
	err = store.db.GetContext(ctx, &result, sql, args...)

	return result, err
}

func articlePageAndKeywordSQL(req request.PaginationRequest) (querySQL, countSQL string, args []interface{}, err error) {
	sql := baseQuerySQLBuilder("cms_article").Where(sq.Eq{"status": "0"}).Where(sq.Eq{"del_flag": "N"})
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		// 根据实际情况填充
		//sql = sql.Where(sq.Or{
		//	sq.Like{"config_name": fmt.Sprint("%", req.Keyword, "%")},
		//})
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

func (store *SQLStore) ArticlePage(ctx context.Context, req request.PaginationRequest) (int64, []internal.CmsArticle, error) {
	var result []internal.CmsArticle
	var total int64

	fail := func(err error) (int64, []internal.CmsArticle, error) {
		return 0, nil, err
	}

	pageSQL, countSQL, args, err := articlePageAndKeywordSQL(req)
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

func (store *SQLStore) ArticleList(ctx context.Context) ([]internal.CmsArticle, error) {
	var result []internal.CmsArticle

	sql, args, err := SQLBuilder().Select("*").From("cms_article").Where(sq.And{
		sq.Eq{"status": "0"},
		sq.Eq{"del_flag": "N"},
	}).ToSql()
	if err != nil {
		return nil, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)

	return result, err
}

func (store *SQLStore) ArticleListByIds(ctx context.Context, ids string) ([]internal.CmsArticle, error) {
	var result []internal.CmsArticle

	sql, args, err := SQLBuilder().Select("*").From("cms_article").
		Where("id = ANY(STRING_TO_ARRAY(?, ',')::int8[])", ids).
		ToSql()
	if err != nil {
		return result, err
	}

	err = store.db.SelectContext(ctx, &result, sql, args...)
	return result, err
}

// ArticleCountByKey 根据关键字统计条数（根据情况启用）
//func (store *SQLStore) ArticleCountByKey(ctx context.Context, key string) (int64, error) {
//	sql, args, err := SQLBuilder().Select("count(*)").From("cms_article").Where(sq.Eq{"ArticleKey": key}).ToSql()
//	if err != nil {
//		return 0, err
//	}
//	var total int64
//	err = store.db.GetContext(ctx, &total, sql, args...)
//	return total, err
//}
