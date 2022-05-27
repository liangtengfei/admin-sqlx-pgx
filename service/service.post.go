package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/model"
	"study.com/demo-sqlx-pgx/utils"
)

func PostList(ctx *gin.Context) ([]model.PostResponse, error) {
	var result []model.PostResponse

	res, err := store.PostList(ctx)
	if err != nil {
		global.Log.Error(BizTitlePost, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func PostPage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.PostResponse, error) {
	var result []model.PostResponse

	total, res, err := store.PostPage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitlePost, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func PostCreate(ctx *gin.Context, req request.PostCreateRequest, username string) error {
	res, err := store.PostCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitlePost, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func PostUpdate(ctx *gin.Context, req request.PostUpdateRequest, username string) error {
	res, err := store.PostUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitlePost, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func PostDeleteFake(ctx *gin.Context, id int64, username string) error {
	res, err := store.PostDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error(BizTitlePost, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func PostDelete(ctx *gin.Context, id int64) error {
	res, err := store.PostDelete(ctx, id)
	if err != nil {
		global.Log.Error(BizTitlePost, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func PostDetail(ctx *gin.Context, id int64) (model.PostResponse, error) {
	var result model.PostResponse

	res, err := store.PostDetail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRows
		}
		global.Log.Error(BizTitlePost, zap.String("TAG", OperationTypeDetail), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func PostListByIds(ctx *gin.Context, ids string) ([]model.PostResponse, error) {
	var result []model.PostResponse

	res, err := store.PostListByIds(ctx, ids)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRows
		}
		global.Log.Error(BizTitlePost, zap.String("TAG", OperationTypeQuery), zap.Error(err))
		return result, ErrQuery
	}
	err = utils.StructCopy(&result, &res)
	return result, err
}
