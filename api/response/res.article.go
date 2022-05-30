package response

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

type ArticleResponse struct {
	Id             int64  `json:"id"`
	ArticleTitle   string `json:"articleTitle"`
	ArticleContent string `json:"articleContent"`
	ArticleTag     string `json:"articleTag"`
	Summary        string `json:"summary"`
	Status         string `json:"status"`
	DelFlag        string `json:"delFlag"`
	CreateTimeStr  string `json:"createTime"`
	UpdateTimeStr  string `json:"updateTime"`
	CreateBy       string `json:"createBy"`
	UpdateBy       string `json:"updateBy"`
	Remark         string `json:"remark"`
}

func (res *ArticleResponse) CreateTime(t time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(t)
}

func (res *ArticleResponse) UpdateTime(t time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(t)
}
