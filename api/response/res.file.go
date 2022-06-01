package response

import (
	"github.com/google/uuid"
	"study.com/demo-sqlx-pgx/utils/datetime"
	"time"
)

type FileResponse struct {
	Id            uuid.UUID `json:"id"`
	FileName      string    `json:"fileName"`
	FilePath      string    `json:"filePath"`
	FileUrl       string    `json:"fileUrl"`
	FileSize      int64     `json:"fileSize"`
	UserId        int64     `json:"userId"`
	MimeType      string    `json:"mimeType"`
	CreateTimeStr string    `json:"createTime"`
	CreateBy      string    `json:"createBy"`
	Remark        string    `json:"remark"`
}

func (res *FileResponse) CreateTime(t time.Time) {
	res.CreateTimeStr = datetime.ToDatetime(t)
}
