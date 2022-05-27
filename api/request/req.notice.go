package request

type NoticeCreateRequest struct {
	NoticeTitle   string `json:"noticeTitle" binding:"required"`   // 公告标题
	NoticeType    string `json:"noticeType" binding:"required"`    // 公告类型（1通知）
	NoticeContent string `json:"noticeContent" binding:"required"` // 公告内容
	Remark        string `json:"remark" binding:"-"`               // 备注
}

type NoticeUpdateRequest struct {
	ID            int64  `json:"id" binding:"required"`            // 公告ID
	NoticeTitle   string `json:"noticeTitle" binding:"required"`   // 公告标题
	NoticeType    string `json:"noticeType" binding:"required"`    // 公告类型（1通知）
	NoticeContent string `json:"noticeContent" binding:"required"` // 公告内容
	Status        string `json:"status" binding:"-"`               // 状态（0正常 1停用）
	Remark        string `json:"remark" binding:"-"`               // 备注
}
