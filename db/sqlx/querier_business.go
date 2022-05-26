package db

import (
	"context"
	"study.com/demo-sqlx-pgx/api/request"
)

// QuerierBusiness 业务接口 使用通知公告举例
type QuerierBusiness interface {
	NoticeCreate(ctx context.Context, req request.NoticeCreateRequest, username string) (int64, error)
	NoticeUpdate(ctx context.Context, req request.NoticeUpdateRequest, username string) (int64, error)
	NoticeDelete(ctx context.Context, id int64) (int64, error)
	NoticeDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	NoticeDetail(ctx context.Context, id int64) (AgoNotice, error)
	NoticePage(ctx context.Context, req request.PaginationRequest) (int64, []AgoNotice, error)
	NoticeList(ctx context.Context) ([]AgoNotice, error)
}
