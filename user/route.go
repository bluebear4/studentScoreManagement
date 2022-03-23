package user

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/util"
)

func Register(ctx *gin.Context) {
	req := &CreateUserRequest{}

	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}

	response := server.CreateUser(ctx, req)

	if response.Base.Code == consts.ErrCodeSuccess {
		//注册后自动登录
		util.Login(ctx, req.ID)
	}
	ctx.JSON(response.Base.ChangeToGinJson(gin.H{"UserID": response.UserID}))
}

func Login(ctx *gin.Context) {
	req := &ValidatePasswordRequest{}

	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}

	response := server.ValidatePassword(ctx, req)
	if response.Base.Code == consts.ErrCodeSuccess {
		util.Login(ctx, req.ID)
	}
	ctx.JSON(response.Base.ChangeToGinJson())

}

func ChangePassword(ctx *gin.Context) {
	req := &ChangePasswordRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	firstResponse := server.ValidatePassword(ctx, &ValidatePasswordRequest{
		ID:       req.ID,
		PassWord: req.OldPassWord,
	})

	if firstResponse.Base.Code == consts.ErrCodeSuccess {
		//校验旧密码
		secondResponse := server.ChangePassword(ctx, req)
		if secondResponse.Base.Code == consts.ErrCodeSuccess {
			util.Login(ctx, req.ID)
		}
		ctx.JSON(secondResponse.Base.ChangeToGinJson())
		return
	}

	ctx.JSON(firstResponse.Base.ChangeToGinJson())
}

func Logout(ctx *gin.Context) {
	//交给鉴权
	util.Logout(ctx)
	ctx.JSON(util.NewBase(consts.ErrCodeSuccess).ChangeToGinJson())
}
