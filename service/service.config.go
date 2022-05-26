package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/model"
	"study.com/demo-sqlx-pgx/utils"
)

func ConfigList(ctx *gin.Context) ([]model.SysConfigResponse, error) {
	var result []model.SysConfigResponse

	res, err := store.ConfigList(ctx)
	if err != nil {
		global.Log.Error(BizTitleConfig, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func ConfigPage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.SysConfigResponse, error) {
	var result []model.SysConfigResponse

	total, res, err := store.ConfigPage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleConfig, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func ConfigCreate(ctx *gin.Context, req request.SysConfigCreateRequest, username string) error {
	res, err := store.ConfigCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleConfig, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func ConfigUpdate(ctx *gin.Context, req request.SysConfigUpdateRequest, username string) error {
	res, err := store.ConfigUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleConfig, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func ConfigDeleteFake(ctx *gin.Context, id int64, username string) error {
	res, err := store.ConfigDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error(BizTitleConfig, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func ConfigDelete(ctx *gin.Context, id int64) error {
	res, err := store.ConfigDelete(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleConfig, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func ConfigDetail(ctx *gin.Context, id int64) (model.SysConfigResponse, error) {
	var result model.SysConfigResponse

	res, err := store.ConfigDetail(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleConfig, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func ConfigBatchCreate(ctx *gin.Context, req []request.SysConfigCreateRequest, username string) (int64, error) {
	res, err := store.ConfigCreateBatch(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleConfig, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return 0, ErrCreate
	}
	return res, nil
}
