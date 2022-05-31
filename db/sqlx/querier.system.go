package db

import (
	"context"
	"github.com/google/uuid"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/db/sqlx/internal"
)

type Querier interface {
	UserFindByUsername(username string) (internal.AgoUser, error)
	UserFindByMobile(mobile string) (internal.AgoUser, error)
	UserCountByMobile(mobile string) (int64, error)
	UserPageAndKeyword(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoUser, error)
	UserDetail(req request.ByIdRequest) (internal.AgoUser, error)
	UserDeleteId(id int64, username string) error
	UserCreate(ctx context.Context, req request.UserCreateRequest, username string) (int64, error)
	UserUpdate(ctx context.Context, req request.UserUpdateRequest, username string) (int64, error)

	RoleListByUserId(id int64) ([]internal.AgoRole, error)
	RoleCreate(ctx context.Context, req request.RoleCreateRequest, username string) (int64, error)
	RoleUpdate(ctx context.Context, req request.RoleUpdateRequest, username string) (int64, error)
	RoleDelete(ctx context.Context, id int64) (int64, error)
	RoleDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	RoleDetail(ctx context.Context, id int64) (internal.AgoRole, error)
	RolePage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoRole, error)
	RoleCountByKey(ctx context.Context, roleKey string) (int64, error)

	MenuListByRoleId(ctx context.Context, id int64) ([]internal.AgoMenu, error)
	MenuListByRoleIds(ctx context.Context, id []int64) ([]internal.AgoMenu, error)
	MenuCreate(ctx context.Context, req request.MenuCreateRequest, username string) (int64, error)
	MenuUpdate(ctx context.Context, req request.MenuUpdateRequest, username string) (int64, error)
	MenuDelete(ctx context.Context, id int64) (int64, error)
	MenuDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	MenuDetail(ctx context.Context, id int64) (internal.AgoMenu, error)
	MenuPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoMenu, error)
	MenuList(ctx context.Context) ([]internal.AgoMenu, error)
	MenuListByIds(ctx context.Context, ids []int64) ([]internal.AgoMenu, error)

	DeptListByRoleId(ctx context.Context, id int64) ([]internal.AgoDept, error)
	DeptCreate(ctx context.Context, req request.DeptCreateRequest, username string) (int64, error)
	DeptUpdate(ctx context.Context, req request.DeptUpdateRequest, username string) (int64, error)
	DeptDelete(ctx context.Context, id int64) (int64, error)
	DeptDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	DeptDetail(ctx context.Context, id int64) (internal.AgoDept, error)
	DeptPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoDept, error)
	DeptList(ctx context.Context) ([]internal.AgoDept, error)

	PostListByIds(ctx context.Context, ids string) ([]internal.AgoPost, error)
	PostCreate(ctx context.Context, req request.PostCreateRequest, username string) (int64, error)
	PostUpdate(ctx context.Context, req request.PostUpdateRequest, username string) (int64, error)
	PostDelete(ctx context.Context, id int64) (int64, error)
	PostDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	PostDetail(ctx context.Context, id int64) (internal.AgoPost, error)
	PostPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoPost, error)
	PostList(ctx context.Context) ([]internal.AgoPost, error)

	ConfigListByIds(ctx context.Context, ids string) ([]internal.AgoConfig, error)
	ConfigCreate(ctx context.Context, req request.SysConfigCreateRequest, username string) (int64, error)
	ConfigCreateBatch(ctx context.Context, req []request.SysConfigCreateRequest, username string) (int64, error)
	ConfigUpdate(ctx context.Context, req request.SysConfigUpdateRequest, username string) (int64, error)
	ConfigDelete(ctx context.Context, id int64) (int64, error)
	ConfigDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	ConfigDetail(ctx context.Context, id int64) (internal.AgoConfig, error)
	ConfigPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoConfig, error)
	ConfigList(ctx context.Context) ([]internal.AgoConfig, error)
	ConfigCountByKey(ctx context.Context, key string) (int64, error)

	DictTypeListByIds(ctx context.Context, ids string) ([]internal.AgoDictType, error)
	DictTypeCreate(ctx context.Context, req request.DictTypeCreateRequest, username string) (int64, error)
	DictTypeUpdate(ctx context.Context, req request.DictTypeUpdateRequest, username string) (int64, error)
	DictTypeDelete(ctx context.Context, id int64) (int64, error)
	DictTypeDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	DictTypeDetail(ctx context.Context, id int64) (internal.AgoDictType, error)
	DictTypePage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoDictType, error)
	DictTypeList(ctx context.Context) ([]internal.AgoDictType, error)

	DictDataListByIds(ctx context.Context, ids string) ([]internal.AgoDictData, error)
	DictDataCreate(ctx context.Context, req request.DictDataCreateRequest, username string) (int64, error)
	DictDataUpdate(ctx context.Context, req request.DictDataUpdateRequest, username string) (int64, error)
	DictDataDelete(ctx context.Context, id int64) (int64, error)
	DictDataDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	DictDataDetail(ctx context.Context, id int64) (internal.AgoDictData, error)
	DictDataPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoDictData, error)
	DictDataList(ctx context.Context) ([]internal.AgoDictData, error)

	SessionListByIds(ctx context.Context, ids string) ([]internal.AgoSession, error)
	SessionCreate(ctx context.Context, req request.SessionCreateRequest, username string) (uuid.UUID, error)
	SessionUpdate(ctx context.Context, req request.SessionUpdateRequest, username string) (int64, error)
	SessionDelete(ctx context.Context, id int64) (int64, error)
	SessionDetail(ctx context.Context, id uuid.UUID) (internal.AgoSession, error)
	SessionPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoSession, error)
	SessionList(ctx context.Context) ([]internal.AgoSession, error)

	OperationLogCreate(ctx context.Context, req request.OperationLogCreateRequest, username string) (int64, error)
	OperationLogPage(ctx context.Context, req request.PaginationRequest) (int64, []internal.AgoOperationLog, error)

	QuerierBusiness
}

var _ Querier = (*SQLStore)(nil)
