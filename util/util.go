package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"studentScoreManagement/config"
	"studentScoreManagement/consts"
	"studentScoreManagement/redis"
	"time"
)

func GetUserID(ctx *gin.Context) (string, error) {
	if cookie, err := ctx.Cookie(consts.CookieAuth); err != nil {
		return "", err
	} else {
		return redis.Get(cookie)
	}
}

func Login(ctx *gin.Context, id string) {
	//登出其他
	Logout(ctx)

	uuid := uuid.New().String()
	ctx.SetCookie(consts.CookieAuth, uuid, 3600*24, "/", config.GetServer().Host, false, false)
	_ = redis.Set(uuid, id, 24*time.Hour)
	_ = redis.Set(id, uuid, 24*time.Hour)
}

func Logout(ctx *gin.Context) {
	if id, err := ctx.Cookie(consts.CookieAuth); err == nil {
		if uuid, err := redis.Get(id); err == nil {
			_ = redis.Del(id)
			_ = redis.Del(uuid)
		}
	}
}
