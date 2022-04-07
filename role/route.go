package role

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/util"
)

func GetValidateCode(ctx *gin.Context) {
	req := &GetValidateCodeRequest{}

	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	response := server.GetValidateCode(ctx, req)
	ctx.JSON(response.Base.ChangeToGinJson(gin.H{
		"ValidateCode": response.ValidateCode,
	}))
}

func ChangeValidateCode(ctx *gin.Context) {
	req := &ChangeValidateCodeRequest{}

	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.ChangeValidateCode(ctx, req).Base.ChangeToGinJson())
}
