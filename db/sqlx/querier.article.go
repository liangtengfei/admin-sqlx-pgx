package db

import (
	"context"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/db/sqlx/internal"
)

// QuerierArticle 业务接口
type QuerierArticle interface {
	ArticleCreate(ctx context.Context, req request.ArticleCreateRequest, username string) (int64, error)
	ArticleCreateBatch(ctx context.Context, req []request.ArticleCreateRequest, username string) (int64, error)
	ArticleUpdate(ctx context.Context, req request.ArticleUpdateRequest, username string) (int64, error)
	ArticleDelete(ctx context.Context, id int64) (int64, error)
	ArticleDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	ArticleDetail(ctx context.Context, id int64) (internal.CmsArticle, error)
	ArticlePage(ctx context.Context, req request.PaginationRequest) (int64, []internal.CmsArticle, error)
	ArticleList(ctx context.Context) ([]internal.CmsArticle, error)
	ArticleListByIds(ctx context.Context, ids string) ([]internal.CmsArticle, error)
}
