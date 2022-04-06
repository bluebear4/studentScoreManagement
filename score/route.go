package score

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/util"
)

func UploadScore(ctx *gin.Context) {
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

func AddScore(ctx *gin.Context) {
	req := &AddScoreRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.AddScore(ctx, req).Base.ChangeToGinJson())
}

func UpdateScore(ctx *gin.Context) {
	req := &UpdateScoreRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.UpdateScore(ctx, req).Base.ChangeToGinJson())
}

func DeleteScore(ctx *gin.Context) {
	req := &DeleteScoreRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.DeleteScore(ctx, req).Base.ChangeToGinJson())
}

func _(ctx *gin.Context) {
	response := server.GetSubjects(ctx)
	if response.Base.Code == consts.ErrCodeSuccess {
		ctx.JSON(response.Base.ChangeToGinJson(gin.H{
			"subjects": response.Subjects,
		}))
	} else {
		ctx.JSON(response.Base.ChangeToGinJson())
	}
	return
}
func GetScores(ctx *gin.Context) {
	req := &GetScoresByIDRequest{}
	if id, err := util.GetUserID(ctx); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	} else {
		req.ID = id
	}

	response := server.GetScoresByID(ctx, req)
	ctx.JSON(response.Base.ChangeToGinJson(gin.H{
		"scores": response.Scores,
	}))
}

func _(ctx *gin.Context) {
	req := &GetScoresByIDRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}

	response := server.GetScoresByID(ctx, req)
	ctx.JSON(response.Base.ChangeToGinJson(gin.H{
		"scores": response.Scores,
	}))
}

func GetScoresByClass(ctx *gin.Context) {
	req := &GetScoresByClassRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}

	response := server.GetScoresByClass(ctx, req)
	ctx.JSON(response.Base.ChangeToGinJson(gin.H{
		"scores": response.Scores,
	}))
}
