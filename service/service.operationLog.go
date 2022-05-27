package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	model "study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/utils"
)

func OperationLogPage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.OperationLogResponse, error) {
	var result []model.OperationLogResponse

	total, res, err := store.OperationLogPage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleOperationLog, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}
