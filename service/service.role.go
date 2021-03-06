package service

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	model "study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/utils"
)

func RoleListByUserId(ctx *gin.Context, userId int64) ([]model.RoleResponse, error) {
	var result []model.RoleResponse

	res, err := store.RoleListByUserId(userId)
	if err != nil {
		global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func RoleKeysByUserId(ctx *gin.Context, userId int64) ([]string, error) {
	var result []string

	roles, err := RoleListByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		result = append(result, role.RoleKey)
	}

	return result, nil
}

func RolePage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.RoleResponse, error) {
	var result []model.RoleResponse

	total, res, err := store.RolePage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrPage
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func RoleList(ctx *gin.Context) ([]model.RoleResponse, error) {
	var result []model.RoleResponse
	req := request.PaginationRequest{
		PageNum:  1,
		PageSize: 100,
	}
	_, res, err := store.RolePage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func RoleCreate(ctx *gin.Context, req request.RoleCreateRequest, username string) error {
	if roleKeyExist(ctx, req.RoleKey, 0) {
		return errors.New("角色标识不能重复")
	}

	_, err := store.RoleCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}

	//菜单权限
	if len(req.MenuIds) > 0 {
		menus, err := store.MenuListByIds(ctx, req.MenuIds)
		if err != nil {
			global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeQuery), zap.Error(err))
			return ErrQuery
		}
		// 如果是按钮的话 则插入casbin 权限策略
		for _, menu := range menus {
			if menu.MenuType != "A" {
				continue
			}
			if exists := global.Enforcer.HasPolicy(req.RoleKey, menu.Path, menu.ReqMethod); !exists {
				ok, err := global.Enforcer.AddPolicy(req.RoleKey, menu.Path, menu.ReqMethod)
				if err != nil && !ok {
					global.Log.Error(BizTitleCasbin, zap.String("TAG", "权限策略新增"), zap.Error(err))
					return ErrUpdate
				}
			}
		}
	}

	return nil
}

func RoleUpdate(ctx *gin.Context, req request.RoleUpdateRequest, username string) error {
	if roleKeyExist(ctx, req.RoleKey, req.ID) {
		return errors.New("角色标识不能重复")
	}

	_, err := store.RoleUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}

	//菜单权限
	if len(req.MenuIds) > 0 {
		menus, err := store.MenuListByIds(ctx, req.MenuIds)
		if err != nil {
			global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeQuery), zap.Error(err))
			return ErrQuery
		}
		// 如果是按钮的话 则插入casbin 权限策略
		for _, menu := range menus {
			if menu.MenuType != "A" {
				continue
			}
			if exists := global.Enforcer.HasPolicy(req.RoleKey, menu.Path, menu.ReqMethod); !exists {
				ok, err := global.Enforcer.AddPolicy(req.RoleKey, menu.Path, menu.ReqMethod)
				if err != nil && !ok {
					global.Log.Error(BizTitleCasbin, zap.String("TAG", "权限策略新增"), zap.Error(err))
					return ErrUpdate
				}
			}
		}
	}

	return nil
}

func RoleDeleteFake(ctx *gin.Context, id int64, username string) error {
	_, err := store.RoleDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}

	return nil
}

func RoleDetail(ctx *gin.Context, id int64) (model.RoleResponse, error) {
	var result model.RoleResponse

	res, err := store.RoleDetail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRows
		}
		global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeDetail), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	if err != nil {
		return result, err
	}

	menus, err := MenuListByRoleId(ctx, id)
	if err != nil {
		return result, err
	}
	result.MenuList = menus

	depts, err := DeptListByRoleId(ctx, id)
	if err != nil {
		return result, err
	}
	result.DeptList = depts

	return result, err
}

// 是否唯一
func roleKeyExist(ctx *gin.Context, roleKey string, id int64) bool {
	total, err := store.RoleCountByKey(ctx, roleKey)
	if err != nil && err == sql.ErrNoRows {
		return false
	}
	// 查询出现错误 同样禁止数据操作
	if err != nil && err != sql.ErrNoRows {
		global.Log.Error(BizTitleRole, zap.String("TAG", OperationTypeQuery), zap.Error(err))
		return true
	}
	if id > 0 {
		return total > 1
	}
	return total > 0
}
