package response

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// PostResponse 岗位信息表
type PostResponse struct {
	ID            int64  `json:"id"`         // 岗位ID
	PostName      string `json:"postName"`   // 岗位名称
	OrderNum      int32  `json:"orderNum"`   // 显示顺序
	Status        string `json:"status"`     // 状态（0正常 1停用）
	DelFlag       string `json:"delFlag"`    // 删除标志（0代表存在 2代表删除）
	CreateBy      string `json:"createBy"`   // 创建者
	CreateTimeStr string `json:"createTime"` // 创建时间
	UpdateBy      string `json:"updateBy"`   // 更新者
	UpdateTimeStr string `json:"updateTime"` // 更新时间
	Remark        string `json:"remark"`     // 备注
}

func (res *PostResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *PostResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}
