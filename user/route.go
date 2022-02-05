package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"studentScoreManagement/config"
	"studentScoreManagement/consts"
	"studentScoreManagement/redis"
	"studentScoreManagement/util"
	"time"
)

func Route(server *gin.Engine) {
	User := server.Group("/user")
	{
		User.POST("/register", register)
		User.POST("/login", login)
		User.POST("/changePassword", changePassword)
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
		//设置cookie
		if _, err := ctx.Cookie(consts.CookieAuth); err != nil {
			uuid := uuid.New().String()
			ctx.SetCookie(consts.CookieAuth, uuid, -1, "/", config.GetServer().Host, true, true)
			_ = redis.Set(uuid, req.UserName, 1*time.Hour)
		}
	}
	ctx.JSON(response.Base.ChangeToGinJson())

}

func login(ctx *gin.Context) {
	req := &ValidatePasswordRequest{}

	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}

	response := server.ValidatePassword(ctx, req)
	if response.Base.Code == consts.ErrCodeSuccess {
		//设置cookie
		if _, err := ctx.Cookie(consts.CookieAuth); err != nil {
			uuid := uuid.New().String()
			ctx.SetCookie(consts.CookieAuth, uuid, -1, "/", config.GetServer().Host, true, true)
			_ = redis.Set(uuid, req.UserName, 1*time.Hour)
		}
	}
	ctx.JSON(response.Base.ChangeToGinJson())

}

func changePassword(ctx *gin.Context) {
	req := &ChangePasswordRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(util.NewBase(consts.ErrCodeParameter).ChangeToGinJson())
		return
	}
	ctx.JSON(server.ChangePassword(ctx, req).Base.ChangeToGinJson())
}
