package user

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/middleware"
	"studentScoreManagement/util"
)

func Route(server *gin.Engine) {
	user := server.Group("/user")
	{
		user.POST("/register", register)
		user.POST("/login", login)
		user.POST("/logout", middleware.Auth(map[int]bool{
			consts.RoleIDStudent: true,
			consts.RoleIDTeacher: true,
			consts.RoleIDAdmin:   true,
		}), logout)
		user.POST("/changePassword", changePassword)
	}
}

func register(ctx *gin.Context) {
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
	ctx.JSON(response.Base.ChangeToGinJson(gin.H{"ID": response.UserID}))
}

func login(ctx *gin.Context) {
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

func logout(ctx *gin.Context) {
	//交给鉴权
	util.Logout(ctx)
	ctx.JSON(util.NewBase(consts.ErrCodeSuccess).ChangeToGinJson())
}

func changePassword(ctx *gin.Context) {
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
