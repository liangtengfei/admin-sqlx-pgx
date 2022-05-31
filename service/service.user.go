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

func UserFindByUsername(ctx *gin.Context, username string) (model.UserResponse, error) {
	res, err := store.UserFindByUsername(username)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "根据用户名查找"), zap.Error(err))
		return model.UserResponse{}, ErrQuery
	}

	var result model.UserResponse
	err = utils.StructCopy(&result, &res)
	if err != nil {
		return result, err
	}

	// 填充角色
	roleKeys, err := RoleKeysByUserId(ctx, res.ID)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联角色"), zap.Error(err))
		return result, ErrQuery
	}
	result.RoleKeys = roleKeys

	roleList, err := RoleListByUserId(ctx, res.ID)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联角色"), zap.Error(err))
		return result, ErrQuery
	}
	result.RoleList = roleList

	return result, err
}

func UserOnlyByUsername(ctx *gin.Context, username string) (model.UserResponse, error) {
	res, err := store.UserFindByUsername(username)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "根据用户名查找"), zap.Error(err))
		return model.UserResponse{}, ErrQuery
	}

	var result model.UserResponse
	err = utils.StructCopy(&result, &res)
	if err != nil {
		return result, err
	}

	return result, err
}

func UserFindByMobile(ctx *gin.Context, username string) (model.UserResponse, error) {
	res, err := store.UserFindByMobile(username)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "根据手机号查找"), zap.Error(err))
		return model.UserResponse{}, ErrQuery
	}

	var result model.UserResponse
	err = utils.StructCopy(&result, &res)
	if err != nil {
		return result, err
	}

	// 填充角色
	roleKeys, err := RoleKeysByUserId(ctx, res.ID)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联角色"), zap.Error(err))
		return result, ErrQuery
	}
	result.RoleKeys = roleKeys

	roleList, err := RoleListByUserId(ctx, res.ID)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联角色"), zap.Error(err))
		return result, ErrQuery
	}
	result.RoleList = roleList

	return result, err
}

func UserDetail(ctx *gin.Context, id int64) (model.UserResponse, error) {
	res, err := store.UserDetail(request.ByIdRequest{Id: id})
	if err != nil {
		if err == sql.ErrNoRows {
			return model.UserResponse{}, ErrNoRows
		}
		global.Log.Error(BizTitleUser, zap.String("TAG", OperationTypeDetail), zap.Error(err))
		return model.UserResponse{}, ErrQuery
	}

	var result model.UserResponse
	err = utils.StructCopy(&result, &res)
	if err != nil {
		return result, err
	}

	// 填充角色
	roleKeys, err := RoleKeysByUserId(ctx, res.ID)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联角色"), zap.Error(err))
		return result, ErrQuery
	}
	result.RoleKeys = roleKeys

	return result, err
}

// UserDetail2 关联其他数据
func UserDetail2(ctx *gin.Context, id int64) (model.UserResponse, error) {
	res, err := store.UserDetail(request.ByIdRequest{Id: id})
	if err != nil {
		if err == sql.ErrNoRows {
			return model.UserResponse{}, ErrNoRows
		}
		global.Log.Error(BizTitleUser, zap.String("TAG", OperationTypeDetail), zap.Error(err))
		return model.UserResponse{}, ErrQuery
	}

	var result model.UserResponse
	err = utils.StructCopy(&result, &res)
	if err != nil {
		return result, err
	}

	// 填充角色
	roleKeys, err := RoleKeysByUserId(ctx, res.ID)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联角色"), zap.Error(err))
		return result, ErrQuery
	}
	result.RoleKeys = roleKeys

	roleList, err := RoleListByUserId(ctx, res.ID)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联角色"), zap.Error(err))
		return result, ErrQuery
	}
	result.RoleList = roleList

	//填充部门
	dept, err := DeptDetail(ctx, res.DeptID)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联部门"), zap.Error(err))
		return result, ErrQuery
	}
	result.Dept = dept

	// 填充岗位
	postList, err := PostListByIds(ctx, res.Posts)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", "查询关联岗位"), zap.Error(err))
		return result, ErrQuery
	}
	result.PostList = postList

	return result, err
}

func UserCreate(ctx *gin.Context, req request.UserCreateRequest, createBy string) (int64, error) {
	var lastId int64 = 0

	password, err := utils.HashPassword(req.Password + "_" + req.Mobile)
	if err != nil {
		return lastId, err
	}
	req.Password = password

	//手机号码是否重复
	err = UniqueMobile(ctx, req.Mobile, 0)
	if err != nil {
		return 0, err
	}

	lastId, err = store.UserCreate(ctx, req, createBy)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return lastId, ErrCreate
	}
	if lastId <= 0 {
		return 0, ErrCreate
	}

	return lastId, nil
}

func UniqueMobile(ctx *gin.Context, mobile string, id int64) error {
	total, err := store.UserCountByMobile(mobile)
	if err == nil && total == 0 {
		return nil
	} else if err != nil && err == sql.ErrNoRows {
		return nil
	} else if id > 0 && total <= 1 {
		return nil
	}
	return errors.New("手机号码已存在")
}

func UserUpdate(ctx *gin.Context, req request.UserUpdateRequest, updateBy string) error {
	//手机号码是否重复
	err := UniqueMobile(ctx, req.Mobile, req.ID)
	if err != nil {
		return err
	}

	rows, err := store.UserUpdate(ctx, req, updateBy)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if rows <= 0 {
		return ErrUpdate
	}
	return nil
}

func UserPageAndKeyword(ctx *gin.Context, req request.PaginationRequest) (int64, []model.UserResponse, error) {
	var result []model.UserResponse

	total, res, err := store.UserPageAndKeyword(ctx, req)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil, ErrNoRows
		}
		global.Log.Error(BizTitleUser, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, nil, ErrPage
	}

	err = utils.StructCopy(&result, &res)

	return total, result, err
}

func UserDeleteById(ctx *gin.Context, id int64, username string) error {
	err := store.UserDeleteId(id, username)
	if err != nil {
		global.Log.Error(BizTitleUser, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}

	return nil
}
