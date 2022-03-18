package middleware

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/model"
	"studentScoreManagement/util"
)

func Auth(roleIDs map[int]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if id, err := util.GetUserID(c); err != nil {
			//未登录 拦截
			c.Abort()
			c.JSON(util.NewBase(consts.ErrCodeNotLogin).ChangeToGinJson())
			return
		} else if len(roleIDs) > 0 {
			r := &model.UserID2RoleID{UserID: id}
			//信任redis
			_ = r.Find()
			if _, ok := roleIDs[r.RoleID]; !ok {
				//权限不够 拦截
				c.Abort()
				c.JSON(util.NewBase(consts.ErrCodeErrorAuthFail).ChangeToGinJson())
				return
			}
		}
		//请求处理
		c.Next()
	}
}
