package scoreView

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/middleware"
)

func Route(server *gin.Engine) {
	score := server.Group("/score", middleware.Auth(map[int]bool{
		consts.RoleIDStudent: true,
		consts.RoleIDTeacher: true,
		consts.RoleIDAdmin:   true,
	}))
	{
		score.GET("/getClass", getClass)
	}
}

func getClass(ctx *gin.Context) {
	response := server.GetClass(ctx)
	if response.Base.Code == consts.ErrCodeSuccess {
		ctx.JSON(response.Base.ChangeToGinJson(gin.H{
			"classes": response.Classes,
		}))
	} else {
		ctx.JSON(response.Base.ChangeToGinJson())
	}
	return
}
