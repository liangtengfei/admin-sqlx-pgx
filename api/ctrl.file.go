package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"path/filepath"
	"study.com/demo-sqlx-pgx/api/request"
	"study.com/demo-sqlx-pgx/api/response"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/service"
)

// FileUploadCommon godoc
// @Summary      文件上传
// @Description  文件上传信息
// @Tags         文件上传
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.FileResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /file/upload/common [post]
func FileUploadCommon(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["files[]"]

	user := GetLoginUserInfo(ctx)
	var reqs []request.FileCreateRequest
	for _, file := range files {
		fileId, err := uuid.NewRandom()
		if err != nil {
			response.ErrorMsg(ctx, "生成UUID错误")
			return
		}
		ext := filepath.Ext(file.Filename)
		if ext == "" {
			response.ErrorMsg(ctx, "获取文件扩展名称错误")
			return
		}
		filePath := fmt.Sprintf("%s/%s%s", global.Config.Server.UploadPath, fileId, ext)

		err = ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			response.ErrorMsg(ctx, err.Error())
			return
		}

		kind, err := filetype.MatchFile(filePath)
		if err != nil || kind == filetype.Unknown {
			response.ErrorMsg(ctx, "获取文件类型错误："+err.Error())
			return
		}
		req := request.FileCreateRequest{
			Id:       fileId,
			FileName: file.Filename,
			FilePath: filePath,
			FileUrl:  "/" + filePath,
			FileSize: file.Size,
			UserId:   user.ID,
			MimeType: kind.MIME.Value,
			Remark:   "无",
		}
		reqs = append(reqs, req)
	}

	err := service.FileCreateBatch(ctx, reqs, user.UserName)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessData(ctx, reqs)
}

// FilePage godoc
// @Summary      文件上传分页查询
// @Description  分页获取所有文件上传信息
// @Tags         文件上传
// @Accept       json
// @Produce      json
// @Success      200       {object}  response.RestRes{data=response.FileResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /file/p [post]
func FilePage(ctx *gin.Context) {
	var req request.PaginationRequest
	if err := ctx.BindQuery(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	total, res, err := service.FilePage(ctx, req)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.SuccessPage(ctx, total, res)
}

// FileDelete godoc
// @Summary      文件上传删除
// @Description  文件上传信息删除
// @Tags         文件上传
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes
// @Failure      500       {object}  response.RestRes
// @Router       /file/:id [delete]
func FileDelete(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}

	err := service.FileDelete(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// FileDetail godoc
// @Summary      文件上传详情
// @Description  文件上传详情信息
// @Tags         文件上传
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "标识"
// @Success      200       {object}  response.RestRes{data=response.FileResponse}
// @Failure      500       {object}  response.RestRes
// @Router       /file/:id [get]
func FileDetail(ctx *gin.Context) {
	var req request.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorValid(ctx, err)
		return
	}
	res, err := service.FileDetail(ctx, req.Id)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, res)
}
