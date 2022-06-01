package internal

import (
	"github.com/google/uuid"
	"time"
)

// AgoFile 上传文件信息表
type AgoFile struct {
	Id         uuid.UUID `json:"id" db:"id"`                  //唯一标识
	FileName   string    `json:"fileName" db:"file_name"`     //上传文件名
	FilePath   string    `json:"filePath" db:"file_path"`     //文件存储路径
	FileUrl    string    `json:"fileUrl" db:"file_url"`       //文件访问路径
	FileSize   int64     `json:"fileSize" db:"file_size"`     //文件的大小
	UserId     int64     `json:"userId" db:"user_id"`         //用户编号
	MimeType   string    `json:"mimeType" db:"mime_type"`     //文件类型
	CreateTime time.Time `json:"createTime" db:"create_time"` //创建时间
	CreateBy   string    `json:"createBy" db:"create_by"`     //创建人员
	Remark     string    `json:"remark" db:"remark"`          //备注
}
