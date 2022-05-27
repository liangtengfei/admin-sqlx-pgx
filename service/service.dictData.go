package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	model "study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/utils"
)

func DictDataList(ctx *gin.Context) ([]model.DictDataResponse, error) {
	var result []model.DictDataResponse

	res, err := store.DictDataList(ctx)
	if err != nil {
		global.Log.Error(BizTitleDictData, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func DictDataPage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.DictDataResponse, error) {
	var result []model.DictDataResponse

	total, res, err := store.DictDataPage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleDictData, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func DictDataCreate(ctx *gin.Context, req request.DictDataCreateRequest, username string) error {
	res, err := store.DictDataCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleDictData, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func DictDataUpdate(ctx *gin.Context, req request.DictDataUpdateRequest, username string) error {
	res, err := store.DictDataUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleDictData, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func DictDataDeleteFake(ctx *gin.Context, id int64, username string) error {
	res, err := store.DictDataDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error(BizTitleDictData, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func DictDataDelete(ctx *gin.Context, id int64) error {
	res, err := store.DictDataDelete(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleDictData, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func DictDataDetail(ctx *gin.Context, id int64) (model.DictDataResponse, error) {
	var result model.DictDataResponse

	res, err := store.DictDataDetail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRows
		}
		global.Log.Error(BizTitleDictData, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}
