package db

import (
	"time"
)

// CmsArticle 文章信息表
type CmsArticle struct {
	Id             int64     `json:"id" db:"id"`                          //唯一标识
	ArticleTitle   string    `json:"articleTitle" db:"article_title"`     //文章标题
	ArticleContent string    `json:"articleContent" db:"article_content"` //文章内容
	ArticleTag     string    `json:"articleTag" db:"article_tag"`         //文章标签
	Summary        string    `json:"summary" db:"summary"`                //摘要
	Status         string    `json:"status" db:"status"`                  //状态（默认0）
	DelFlag        string    `json:"delFlag" db:"del_flag"`               //删除标记
	CreateTime     time.Time `json:"createTime" db:"create_time"`         //创建时间
	UpdateTime     time.Time `json:"updateTime" db:"update_time"`         //更新时间
	CreateBy       string    `json:"createBy" db:"create_by"`             //创建人员
	UpdateBy       string    `json:"updateBy" db:"update_by"`             //更新人员
	Remark         string    `json:"remark" db:"remark"`                  //备注
}
