package request

import "github.com/google/uuid"

type FileCreateRequest struct {
	Id       uuid.UUID `json:"id" binding:"required"`
	FileName string    `json:"fileName" binding:"required"`
	FilePath string    `json:"filePath" binding:"required"`
	FileUrl  string    `json:"fileUrl" binding:"required"`
	FileSize int64     `json:"fileSize" binding:"required"`
	UserId   int64     `json:"userId" binding:"required"`
	MimeType string    `json:"mimeType" binding:"required"`
	Remark   string    `json:"remark" binding:"-"`
}

type FileUpdateRequest struct {
	Id       uuid.UUID `json:"id" binding:"required"`
	FileName string    `json:"fileName" binding:"required"`
	FilePath string    `json:"filePath" binding:"required"`
	FileUrl  string    `json:"fileUrl" binding:"required"`
	FileSize int64     `json:"fileSize" binding:"required"`
	UserId   int64     `json:"userId" binding:"required"`
	MimeType string    `json:"mimeType" binding:"required"`
	Remark   string    `json:"remark" binding:"-"`
}

type FileIdsRequest struct {
	IdList []int64 `json:"ids" form:"ids"`
}
