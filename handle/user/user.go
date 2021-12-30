package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"studentScoreManagement/config"
	"studentScoreManagement/handle"
	"studentScoreManagement/model"
	"studentScoreManagement/redis"
	"time"
)

const (
	userCookie = "session"
)

func GetUserID(c *gin.Context) (string, error) {
	if cookie, err := c.Cookie(userCookie); err != nil {
		return "", err
	} else {
		return redis.Get(cookie)
	}
}

func Register(c *gin.Context) {
	var form handle.RegisterFrom
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	} else {
		role := &model.Role{RoleID: form.RoleID}
		if err := role.Find(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
			return
		} else if role.RoleID == "1" && role.RoleCode != "" && role.RoleCode != form.VerifyCode {
			//1:教师 为空时禁止注册
			c.JSON(http.StatusBadRequest, gin.H{"msg": "验证码错误"})
			return
		}
	}

	//用户注册
	user := &model.User{
		Name:     form.UserName,
		PassWord: form.PassWord,
	}
	if err := user.Create(); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名已存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误,稍后重试"})
		return
	}

	//角色绑定
	role := &model.UserID2RoleID{
		UserID: user.ID,
		RoleID: form.RoleID,
	}
	if err := role.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误,稍后重试"})
		return
	}

	//设置cookie
	if _, err := c.Cookie(userCookie); err != nil {
		uuid := uuid.New().String()
		c.SetCookie(userCookie, uuid, -1, "/", config.GetServer().Host, true, true)
		_ = redis.Set(uuid, form.UserName, 1*time.Hour)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

func Login(c *gin.Context) {
	var form handle.LoginFrom
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	user := &model.User{
		Name:     form.UserName,
		PassWord: form.PassWord,
	}
	if err := user.Find(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名或密码错误"})
		return
	}

	//设置cookie
	if _, err := c.Cookie(userCookie); err != nil {
		uuid := uuid.New().String()
		c.SetCookie(userCookie, uuid, -1, "/", config.GetServer().Host, true, true)
		_ = redis.Set(uuid, form.UserName, 1*time.Hour)
	}

	c.Redirect(http.StatusTemporaryRedirect, "/index")
}
