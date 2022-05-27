package response

import (
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

// NoticeResponse 通知公告信息
type NoticeResponse struct {
	ID            int64  `json:"id"`            // 公告ID
	NoticeTitle   string `json:"noticeTitle"`   // 公告标题
	NoticeType    string `json:"noticeType"`    // 公告类型（1通知）
	NoticeContent string `json:"noticeContent"` // 公告内容
	Status        string `json:"status"`        // 状态（0正常 1停用）
	DelFlag       string `json:"delFlag"`       // 删除标志（0代表存在 2代表删除）
	CreateBy      string `json:"createBy"`      // 创建者
	CreateTimeStr string `json:"createTime"`    // 创建时间
	UpdateBy      string `json:"updateBy"`      // 更新者
	UpdateTimeStr string `json:"updateTime"`    // 更新时间
	Remark        string `json:"remark"`        // 备注
}

func (res *NoticeResponse) CreateTime(createTime time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(createTime)
}

func (res *NoticeResponse) UpdateTime(updateTime time.Time) {
	res.UpdateTimeStr = datetime.ToDatetime(updateTime)
}
