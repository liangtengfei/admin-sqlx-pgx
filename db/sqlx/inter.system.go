package db

import (
	"context"
	"github.com/google/uuid"
	"study.com/demo-sqlx-pgx/api/request"
)

type Querier interface {
	UserFindByUsername(username string) (AgoUser, error)
	UserFindByMobile(mobile string) (AgoUser, error)
	UserCountByMobile(mobile string) (int64, error)
	UserPageAndKeyword(ctx context.Context, req request.PaginationRequest) (int64, []AgoUser, error)
	UserDetail(req request.ByIdRequest) (AgoUser, error)
	UserDeleteId(id int64, username string) error
	UserCreate(ctx context.Context, req request.UserCreateRequest, username string) (int64, error)
	UserUpdate(ctx context.Context, req request.UserUpdateRequest, username string) (int64, error)

	RoleListByUserId(id int64) ([]AgoRole, error)
	RoleCreate(ctx context.Context, req request.RoleCreateRequest, username string) (int64, error)
	RoleUpdate(ctx context.Context, req request.RoleUpdateRequest, username string) (int64, error)
	RoleDelete(ctx context.Context, id int64) (int64, error)
	RoleDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	RoleDetail(ctx context.Context, id int64) (AgoRole, error)
	RolePage(ctx context.Context, req request.PaginationRequest) (int64, []AgoRole, error)
	RoleCountByKey(ctx context.Context, roleKey string) (int64, error)

	MenuListByRoleId(ctx context.Context, id int64) ([]AgoMenu, error)
	MenuCreate(ctx context.Context, req request.MenuCreateRequest, username string) (int64, error)
	MenuUpdate(ctx context.Context, req request.MenuUpdateRequest, username string) (int64, error)
	MenuDelete(ctx context.Context, id int64) (int64, error)
	MenuDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	MenuDetail(ctx context.Context, id int64) (AgoMenu, error)
	MenuPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoMenu, error)
	MenuList(ctx context.Context) ([]AgoMenu, error)

	DeptListByRoleId(ctx context.Context, id int64) ([]AgoDept, error)
	DeptCreate(ctx context.Context, req request.DeptCreateRequest, username string) (int64, error)
	DeptUpdate(ctx context.Context, req request.DeptUpdateRequest, username string) (int64, error)
	DeptDelete(ctx context.Context, id int64) (int64, error)
	DeptDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	DeptDetail(ctx context.Context, id int64) (AgoDept, error)
	DeptPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoDept, error)
	DeptList(ctx context.Context) ([]AgoDept, error)

	PostListByIds(ctx context.Context, ids string) ([]AgoPost, error)
	PostCreate(ctx context.Context, req request.PostCreateRequest, username string) (int64, error)
	PostUpdate(ctx context.Context, req request.PostUpdateRequest, username string) (int64, error)
	PostDelete(ctx context.Context, id int64) (int64, error)
	PostDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	PostDetail(ctx context.Context, id int64) (AgoPost, error)
	PostPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoPost, error)
	PostList(ctx context.Context) ([]AgoPost, error)

	ConfigListByIds(ctx context.Context, ids string) ([]AgoConfig, error)
	ConfigCreate(ctx context.Context, req request.SysConfigCreateRequest, username string) (int64, error)
	ConfigCreateBatch(ctx context.Context, req []request.SysConfigCreateRequest, username string) (int64, error)
	ConfigUpdate(ctx context.Context, req request.SysConfigUpdateRequest, username string) (int64, error)
	ConfigDelete(ctx context.Context, id int64) (int64, error)
	ConfigDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	ConfigDetail(ctx context.Context, id int64) (AgoConfig, error)
	ConfigPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoConfig, error)
	ConfigList(ctx context.Context) ([]AgoConfig, error)
	ConfigCountByKey(ctx context.Context, key string) (int64, error)

	DictTypeListByIds(ctx context.Context, ids string) ([]AgoDictType, error)
	DictTypeCreate(ctx context.Context, req request.DictTypeCreateRequest, username string) (int64, error)
	DictTypeUpdate(ctx context.Context, req request.DictTypeUpdateRequest, username string) (int64, error)
	DictTypeDelete(ctx context.Context, id int64) (int64, error)
	DictTypeDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	DictTypeDetail(ctx context.Context, id int64) (AgoDictType, error)
	DictTypePage(ctx context.Context, req request.PaginationRequest) (int64, []AgoDictType, error)
	DictTypeList(ctx context.Context) ([]AgoDictType, error)

	DictDataListByIds(ctx context.Context, ids string) ([]AgoDictData, error)
	DictDataCreate(ctx context.Context, req request.DictDataCreateRequest, username string) (int64, error)
	DictDataUpdate(ctx context.Context, req request.DictDataUpdateRequest, username string) (int64, error)
	DictDataDelete(ctx context.Context, id int64) (int64, error)
	DictDataDeleteFake(ctx context.Context, id int64, username string) (int64, error)
	DictDataDetail(ctx context.Context, id int64) (AgoDictData, error)
	DictDataPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoDictData, error)
	DictDataList(ctx context.Context) ([]AgoDictData, error)

	SessionListByIds(ctx context.Context, ids string) ([]AgoSession, error)
	SessionCreate(ctx context.Context, req request.SessionCreateRequest, username string) (uuid.UUID, error)
	SessionUpdate(ctx context.Context, req request.SessionUpdateRequest, username string) (int64, error)
	SessionDelete(ctx context.Context, id int64) (int64, error)
	SessionDetail(ctx context.Context, id uuid.UUID) (AgoSession, error)
	SessionPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoSession, error)
	SessionList(ctx context.Context) ([]AgoSession, error)

	OperationLogCreate(ctx context.Context, req request.OperationLogCreateRequest, username string) (int64, error)
	OperationLogPage(ctx context.Context, req request.PaginationRequest) (int64, []AgoOperationLog, error)

	QuerierBusiness
}

var _ Querier = (*SQLStore)(nil)
