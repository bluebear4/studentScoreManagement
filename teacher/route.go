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

		upload := teacher.Group("/upload")
		{
			upload.POST("/info", uploadInfo)
		}

	}
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
