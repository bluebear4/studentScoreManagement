package teacher

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/middleware"
	"studentScoreManagement/util"
)

func Route(server *gin.Engine) {
	teacher := server.Group("/teacher", middleware.Auth(map[int]bool{
		consts.RoleIDTeacher: true,
		consts.RoleIDAdmin:   true,
	}))
	{
		teacher.POST("/addInfo", addInfo)
		teacher.POST("/updateInfo", updateInfo)
		teacher.POST("/deleteInfo", deleteInfo)

		teacher.POST("/addScore", addScore)
		teacher.POST("/updateScore", updateScore)
		teacher.POST("/deleteScore", deleteScore)

		upload := teacher.Group("/upload")
		{
			upload.POST("/info", uploadInfo)
			upload.POST("/score", uploadScore)
		}
	}
}

func uploadInfo(ctx *gin.Context) {
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

func addInfo(ctx *gin.Context) {
	req := &AddInfoRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.AddInfo(ctx, req).Base.ChangeToGinJson())
}

func updateInfo(ctx *gin.Context) {
	req := &UpdateInfoRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.UpdateInfo(ctx, req).Base.ChangeToGinJson())
}

func deleteInfo(ctx *gin.Context) {
	req := &DeleteInfoRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.DeleteInfo(ctx, req).Base.ChangeToGinJson())
}

func uploadScore(ctx *gin.Context) {
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

	response := server.UploadScore(ctx, &UploadScoreRequest{File: xlsx})
	ctx.JSON(response.Base.ChangeToGinJson(gin.H{
		"SuccessCount": response.SuccessCount,
		"FailCount":    response.FailCount,
	}))
}

func addScore(ctx *gin.Context) {
	req := &AddScoreRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.AddScore(ctx, req).Base.ChangeToGinJson())
}

func updateScore(ctx *gin.Context) {
	req := &UpdateScoreRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.UpdateScore(ctx, req).Base.ChangeToGinJson())
}

func deleteScore(ctx *gin.Context) {
	req := &DeleteScoreRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.DeleteScore(ctx, req).Base.ChangeToGinJson())
}
