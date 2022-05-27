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

func MenuListByRoleId(ctx *gin.Context, id int64) ([]model.MenuResponse, error) {
	var result []model.MenuResponse

	rows, err := store.MenuListByRoleId(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleMenu, zap.String("TAG", OperationTypeQuery), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &rows)
	return result, err
}

func MenuPage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.MenuResponse, error) {
	var result []model.MenuResponse

	total, res, err := store.MenuPage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleMenu, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func MenuCreate(ctx *gin.Context, req request.MenuCreateRequest, username string) error {
	res, err := store.MenuCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleMenu, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func MenuUpdate(ctx *gin.Context, req request.MenuUpdateRequest, username string) error {
	res, err := store.MenuUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleMenu, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func MenuDeleteFake(ctx *gin.Context, id int64, username string) error {
	res, err := store.MenuDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error(BizTitleMenu, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func MenuDelete(ctx *gin.Context, id int64) error {
	res, err := store.MenuDelete(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleMenu, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func MenuDetail(ctx *gin.Context, id int64) (model.MenuResponse, error) {
	var result model.MenuResponse

	res, err := store.MenuDetail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, ErrNoRows
		}
		global.Log.Error("系统部门-详情", zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func MenuListTree(ctx *gin.Context) ([]*model.MenuResponse, error) {
	var result []model.MenuResponse

	res, err := store.MenuList(ctx)
	if err != nil {
		global.Log.Error(BizTitleMenu, zap.String("TAG", OperationTypeQuery), zap.Error(err))
		return nil, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	//整理成树形
	list := make([]*model.MenuResponse, 0)
	for i := 0; i < len(result); i++ {
		list = append(list, &result[i])
	}

	tree := make([]*model.MenuResponse, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentID == 0 {
			makeMenuTree(list, list[i])
			tree = append(tree, list[i])
		}
	}
	return tree, err
}

//查找子列表
func findMenuChild(list []*model.MenuResponse, v *model.MenuResponse) (res []*model.MenuResponse) {
	for _, vfor := range list {
		if vfor.ParentID == v.ID {
			res = append(res, vfor)
		}
	}
	return
}

func hasMenuChild(list []*model.MenuResponse, v *model.MenuResponse) bool {
	return len(findMenuChild(list, v)) > 0
}

//格式化
func makeMenuTree(list []*model.MenuResponse, v *model.MenuResponse) {
	children := findMenuChild(list, v)
	for _, child := range children {
		v.Children = append(v.Children, child)
		if hasMenuChild(list, child) {
			makeMenuTree(list, child)
		}
	}
}
