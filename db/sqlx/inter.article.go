package db

import (
	"context"
    "study.com/demo-sqlx-pgx/api/request"
)
// QuerierArticle 业务接口
type QuerierArticle interface {
    ArticleCreate(ctx context.Context, req request.ArticleCreateRequest, username string) (int64, error)
    ArticleCreateBatch(ctx context.Context, req []request.ArticleCreateRequest, username string) (int64, error)
    ArticleUpdate(ctx context.Context, req request.ArticleUpdateRequest, username string) (int64, error)
    ArticleDelete(ctx context.Context, id int64) (int64, error)
    ArticleDeleteFake(ctx context.Context, id int64, username string) (int64, error)
    ArticleDetail(ctx context.Context, id int64) (CmsArticle, error)
    ArticlePage(ctx context.Context, req request.PaginationRequest) (int64, []CmsArticle, error)
    ArticleList(ctx context.Context) ([]CmsArticle, error)
    ArticleListByIds(ctx context.Context, ids string) ([]CmsArticle, error)
}