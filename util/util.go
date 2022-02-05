package util

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
	"studentScoreManagement/redis"
)

func GetUserID(ctx *gin.Context) (string, error) {
	if cookie, err := ctx.Cookie(consts.CookieAuth); err != nil {
		return "", err
	} else {
		return redis.Get(cookie)
	}
}
