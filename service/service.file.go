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

func FilePage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.FileResponse, error) {
	var result []model.FileResponse

	total, res, err := store.FilePage(ctx, req)
	if err != nil {
		global.Log.Error("File", zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func FileCreate(ctx *gin.Context, req request.FileCreateRequest, username string) error {
	res, err := store.FileCreate(ctx, req, username)
	if err != nil {
		global.Log.Error("File", zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func FileCreateBatch(ctx *gin.Context, req []request.FileCreateRequest, username string) error {
	res, err := store.FileCreateBatch(ctx, req, username)
	if err != nil {
		global.Log.Error("File", zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func FileDelete(ctx *gin.Context, id int64) error {
	res, err := store.FileDelete(ctx, id)
	if err != nil {
		global.Log.Error("File", zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func FileDetail(ctx *gin.Context, id int64) (model.FileResponse, error) {
	var result model.FileResponse

	res, err := store.FileDetail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRows
		}
		global.Log.Error("File", zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}
