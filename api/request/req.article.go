package request

type ArticleCreateRequest struct {
	ArticleTitle   string `json:"articleTitle" binding:"required"`
	ArticleContent string `json:"articleContent" binding:"required"`
	ArticleTag     string `json:"articleTag" binding:"required"`
	Summary        string `json:"summary" binding:"required"`
	Remark         string `json:"remark" binding:"-"`
}

type ArticleUpdateRequest struct {
	Id             int64  `json:"id" binding:"required"`
	ArticleTitle   string `json:"articleTitle" binding:"required"`
	ArticleContent string `json:"articleContent" binding:"required"`
	ArticleTag     string `json:"articleTag" binding:"required"`
	Summary        string `json:"summary" binding:"required"`
	Status         string `json:"status" binding:"required"`
	Remark         string `json:"remark" binding:"-"`
}

type ArticleIdsRequest struct {
	IdList []int64 `json:"ids" form:"ids"`
}
