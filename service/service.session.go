package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	model "study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/utils"
)

func SessionList(ctx *gin.Context) ([]model.SessionResponse, error) {
	var result []model.SessionResponse

	res, err := store.SessionList(ctx)
	if err != nil {
		global.Log.Error(BizTitleSession, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func SessionPage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.SessionResponse, error) {
	var result []model.SessionResponse

	total, res, err := store.SessionPage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleSession, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func SessionCreate(ctx *gin.Context, req request.SessionCreateRequest, username string) (uuid.UUID, error) {
	var id uuid.UUID
	id, err := store.SessionCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleSession, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return id, ErrCreate
	}

	return id, nil
}

func SessionUpdate(ctx *gin.Context, req request.SessionUpdateRequest, username string) error {
	res, err := store.SessionUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleSession, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func SessionDelete(ctx *gin.Context, id int64) error {
	res, err := store.SessionDelete(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleSession, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func SessionDetail(ctx *gin.Context, id uuid.UUID) (model.SessionResponse, error) {
	var result model.SessionResponse

	res, err := store.SessionDetail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRows
		}
		global.Log.Error(BizTitleSession, zap.String("TAG", OperationTypeDetail), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}
