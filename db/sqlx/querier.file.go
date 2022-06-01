package db

import (
	"context"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/db/sqlx/internal"
)

// QuerierFile 业务接口
type QuerierFile interface {
	FileCreate(ctx context.Context, req request.FileCreateRequest, username string) (int64, error)
	FileCreateBatch(ctx context.Context, req []request.FileCreateRequest, username string) (int64, error)
	FileUpdate(ctx context.Context, req request.FileUpdateRequest, username string) (int64, error)
	FileDelete(ctx context.Context, id int64) (int64, error)
	FileDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	FileDetail(ctx context.Context, id int64) (internal.AgoFile, error)
	FilePage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoFile, error)
	FileList(ctx context.Context) ([]internal.AgoFile, error)
	FileListByIds(ctx context.Context, ids string) ([]internal.AgoFile, error)
}
