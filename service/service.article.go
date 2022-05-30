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

func ArticleList(ctx *gin.Context) ([]model.ArticleResponse, error) {
	var result []model.ArticleResponse

	res, err := store.ArticleList(ctx)
	if err != nil {
		global.Log.Error("Article", zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func ArticlePage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.ArticleResponse, error) {
	var result []model.ArticleResponse

	total, res, err := store.ArticlePage(ctx, req)
	if err != nil {
		global.Log.Error("Article", zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func ArticleCreate(ctx *gin.Context, req request.ArticleCreateRequest, username string) error {
	res, err := store.ArticleCreate(ctx, req, username)
	if err != nil {
		global.Log.Error("Article", zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func ArticleUpdate(ctx *gin.Context, req request.ArticleUpdateRequest, username string) error {
	res, err := store.ArticleUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error("Article", zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func ArticleDeleteFake(ctx *gin.Context, id int64, username string) error {
	res, err := store.ArticleDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error("Article", zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func ArticleDelete(ctx *gin.Context, id int64) error {
	res, err := store.ArticleDelete(ctx, id)
	if err != nil {
		global.Log.Error("Article", zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func ArticleDetail(ctx *gin.Context, id int64) (model.ArticleResponse, error) {
	var result model.ArticleResponse

	res, err := store.ArticleDetail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRows
		}
		global.Log.Error("Article", zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func ArticleBatchCreate(ctx *gin.Context, req []request.ArticleCreateRequest, username string) (int64, error) {
	res, err := store.ArticleCreateBatch(ctx, req, username)
	if err != nil {
		global.Log.Error("Article", zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return 0, ErrCreate
	}
	return res, nil
}
