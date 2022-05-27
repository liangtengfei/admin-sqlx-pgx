package request

// PaginationRequest 分页通用请求参数
type PaginationRequest struct {
	PageNum   int32  `json:"pageNum" form:"pageNum" binding:"required,min=1"`   // 页码
	PageSize  int32  `json:"pageSize" form:"pageSize" binding:"required,min=1"` // 每页条数
	Keyword   string `json:"keyword" form:"keyword" binding:"-"`
	SortField string `json:"sortField" form:"sortField"` // 排序字段
	SortOrder string `json:"sortOrder" form:"sortOrder"` // 排序顺序
}

type DataScopeRequest struct {
	Scope  string        `binding:"-"`
	Params []interface{} `binding:"-"`
}

type ByIdRequest struct {
	Id int64 `json:"id" form:"id" uri:"id" binding:"required"`
}

func (req PaginationRequest) GetOffset() uint64 {
	if req.PageNum > 10000 || req.PageNum <= 0 {
		req.PageSize = 1
	}
	var offset int32
	if req.PageSize <= 0 || req.PageSize > 10000 {
		offset = 10
	} else {
		offset = (req.PageNum - 1) * req.PageSize
	}
	return uint64(offset)
}

func (req PaginationRequest) GetLimit() uint64 {
	var limit int32
	if req.PageSize <= 0 || req.PageSize > 10000 {
		limit = 10
	} else {
		limit = req.PageSize
	}
	return uint64(limit)
}
