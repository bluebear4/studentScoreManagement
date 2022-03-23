package info

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/util"
)

func UploadInfo(ctx *gin.Context) {
	// 获取上传文件
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	//读取文件
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}

	response := server.UploadInfo(ctx, &UploadInfoRequest{File: xlsx})
	ctx.JSON(response.Base.ChangeToGinJson(gin.H{
		"SuccessCount": response.SuccessCount,
		"FailCount":    response.FailCount,
	}))
}

func AddInfo(ctx *gin.Context) {
	req := &AddInfoRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.AddInfo(ctx, req).Base.ChangeToGinJson())
}

func UpdateInfo(ctx *gin.Context) {
	req := &UpdateInfoRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.UpdateInfo(ctx, req).Base.ChangeToGinJson())
}

func DeleteInfo(ctx *gin.Context) {
	req := &DeleteInfoRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.DeleteInfo(ctx, req).Base.ChangeToGinJson())
}

func GetClasses(ctx *gin.Context) {
	response := server.GetClasses(ctx)
	if response.Base.Code == consts.ErrCodeSuccess {
		ctx.JSON(response.Base.ChangeToGinJson(gin.H{
			"classes": response.Classes,
		}))
	} else {
		ctx.JSON(response.Base.ChangeToGinJson())
	}
	return
}
