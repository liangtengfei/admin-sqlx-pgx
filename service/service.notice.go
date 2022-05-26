package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/model"
	"study.com/demo-sqlx-pgx/utils"
)

func NoticeList(ctx *gin.Context) ([]model.NoticeResponse, error) {
	var result []model.NoticeResponse

	res, err := store.NoticeList(ctx)
	if err != nil {
		global.Log.Error(BizTitleNotice, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func NoticePage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.NoticeResponse, error) {
	var result []model.NoticeResponse

	total, res, err := store.NoticePage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleNotice, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func NoticeCreate(ctx *gin.Context, req request.NoticeCreateRequest, username string) error {
	res, err := store.NoticeCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleNotice, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func NoticeUpdate(ctx *gin.Context, req request.NoticeUpdateRequest, username string) error {
	res, err := store.NoticeUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleNotice, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func NoticeDeleteFake(ctx *gin.Context, id int64, username string) error {
	res, err := store.NoticeDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error(BizTitleNotice, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func NoticeDelete(ctx *gin.Context, id int64) error {
	res, err := store.NoticeDelete(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleNotice, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func NoticeDetail(ctx *gin.Context, id int64) (model.NoticeResponse, error) {
	var result model.NoticeResponse

	res, err := store.NoticeDetail(ctx, id)
	if err != nil {
		global.Log.Error("系统部门-详情", zap.Error(err))
		return result, err
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}
