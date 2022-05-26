package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/model"
	"study.com/demo-sqlx-pgx/utils"
)

func DictTypeList(ctx *gin.Context) ([]model.DictTypeResponse, error) {
	var result []model.DictTypeResponse

	res, err := store.DictTypeList(ctx)
	if err != nil {
		global.Log.Error(BizTitleDictType, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func DictTypePage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.DictTypeResponse, error) {
	var result []model.DictTypeResponse

	total, res, err := store.DictTypePage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleDictType, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func DictTypeCreate(ctx *gin.Context, req request.DictTypeCreateRequest, username string) error {
	res, err := store.DictTypeCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleDictType, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func DictTypeUpdate(ctx *gin.Context, req request.DictTypeUpdateRequest, username string) error {
	res, err := store.DictTypeUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleDictType, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func DictTypeDeleteFake(ctx *gin.Context, id int64, username string) error {
	res, err := store.DictTypeDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error(BizTitleDictType, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func DictTypeDelete(ctx *gin.Context, id int64) error {
	res, err := store.DictTypeDelete(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleDictType, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func DictTypeDetail(ctx *gin.Context, id int64) (model.DictTypeResponse, error) {
	var result model.DictTypeResponse

	res, err := store.DictTypeDetail(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleDictType, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}
