package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/model"
	"study.com/demo-sqlx-pgx/utils"
)

func DeptList(ctx *gin.Context) ([]model.DeptResponse, error) {
	var result []model.DeptResponse

	res, err := store.DeptList(ctx)
	if err != nil {
		global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeList), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func DeptListByRoleId(ctx *gin.Context, id int64) ([]model.DeptResponse, error) {
	var result []model.DeptResponse

	res, err := store.DeptListByRoleId(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeQuery), zap.Error(err))
		return result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

func DeptListTree(ctx *gin.Context) ([]*model.DeptResponse, error) {
	var result []model.DeptResponse

	res, err := store.DeptList(ctx)
	if err != nil {
		global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeList), zap.Error(err))
		return nil, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	//整理成树形
	list := make([]*model.DeptResponse, 0)
	for i := 0; i < len(result); i++ {
		list = append(list, &result[i])
	}

	tree := make([]*model.DeptResponse, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentID == 0 {
			makeTree(list, list[i])
			tree = append(tree, list[i])
		}
	}
	return tree, err
}

func DeptPage(ctx *gin.Context, req request.PaginationRequest) (int64, []model.DeptResponse, error) {
	var result []model.DeptResponse

	total, res, err := store.DeptPage(ctx, req)
	if err != nil {
		global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypePage), zap.Error(err))
		return 0, result, ErrQuery
	}

	err = utils.StructCopy(&result, &res)
	return total, result, err
}

func DeptCreate(ctx *gin.Context, req request.DeptCreateRequest, username string) error {
	ancestors := "0"
	if req.ParentID > 0 {
		parent, err := store.DeptDetail(ctx, req.ParentID)
		if err != nil {
			global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeDetail), zap.Error(err))
			return ErrQuery
		}
		ancestors = utils.StringConcat(parent.Ancestors, ",", utils.Int2String(parent.ID))
	}
	req.Ancestors = ancestors

	res, err := store.DeptCreate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeCreate), zap.Error(err))
		return ErrCreate
	}
	if res <= 0 {
		return ErrCreate
	}

	return nil
}

func DeptUpdate(ctx *gin.Context, req request.DeptUpdateRequest, username string) error {
	ancestors := "0"
	if req.ParentID > 0 {
		parent, err := store.DeptDetail(ctx, req.ParentID)
		if err != nil {
			global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeDetail), zap.Error(err))
			return ErrQuery
		}
		ancestors = utils.StringConcat(parent.Ancestors, ",", utils.Int2String(parent.ID))
	}
	req.Ancestors = ancestors

	res, err := store.DeptUpdate(ctx, req, username)
	if err != nil {
		global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeUpdate), zap.Error(err))
		return ErrUpdate
	}
	if res <= 0 {
		return ErrUpdate
	}
	return nil
}

func DeptDeleteFake(ctx *gin.Context, id int64, username string) error {
	res, err := store.DeptDeleteFake(ctx, id, username)
	if err != nil {
		global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func DeptDelete(ctx *gin.Context, id int64) error {
	res, err := store.DeptDelete(ctx, id)
	if err != nil {
		global.Log.Error(BizTitleDept, zap.String("TAG", OperationTypeDelete), zap.Error(err))
		return ErrDelete
	}
	if res <= 0 {
		return ErrDelete
	}
	return nil
}

func DeptDetail(ctx *gin.Context, id int64) (model.DeptResponse, error) {
	var result model.DeptResponse

	res, err := store.DeptDetail(ctx, id)
	if err != nil {
		global.Log.Error("系统部门-详情", zap.Error(err))
		return result, err
	}

	err = utils.StructCopy(&result, &res)
	return result, err
}

//查找子列表
func findChild(list []*model.DeptResponse, v *model.DeptResponse) (res []*model.DeptResponse) {
	for _, vfor := range list {
		if vfor.ParentID == v.ID {
			res = append(res, vfor)
		}
	}
	return
}

func hasChild(list []*model.DeptResponse, v *model.DeptResponse) bool {
	return len(findChild(list, v)) > 0
}

//格式化
func makeTree(list []*model.DeptResponse, v *model.DeptResponse) {
	children := findChild(list, v)
	for _, child := range children {
		v.Children = append(v.Children, child)
		if hasChild(list, child) {
			makeTree(list, child)
		}
	}
}
