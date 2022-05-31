package db

import (
	"context"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/db/sqlx/internal"
)

// QuerierNotice 业务接口 使用通知公告举例
type QuerierNotice interface {
	NoticeCreate(ctx context.Context, req request.NoticeCreateRequest, username string) (int64, error)
	NoticeUpdate(ctx context.Context, req request.NoticeUpdateRequest, username string) (int64, error)
	NoticeDelete(ctx context.Context, id int64) (int64, error)
	NoticeDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	NoticeDetail(ctx context.Context, id int64) (internal.AgoNotice, error)
	NoticePage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoNotice, error)
	NoticeList(ctx context.Context) ([]internal.AgoNotice, error)
}
