package middleware

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/util"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := util.GetUserID(c); err != nil {
			//未登录 拦截
			c.JSON(util.NewBase(consts.ErrCodeNotLogin).ChangeToGinJson())
		}

		//请求处理
		c.Next()
	}
}
