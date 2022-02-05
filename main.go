package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"studentScoreManagement/config"
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
	"studentScoreManagement/model"
	"studentScoreManagement/redis"
	"studentScoreManagement/user"
)

func init() {
	config.InitConfig()
	db.InitDatabase()
	err := db.GetDatabase().AutoMigrate(&model.User{}, &model.Role{}, &model.UserID2RoleID{})
	if err != nil {
		panic("建表失败" + err.Error())
	}
	//hard code 设定角色
	role := model.Role{
		RoleID:   consts.RoleIDAdmin,
		RoleName: consts.RoleNameAdmin,
	}
	if err := role.Find(); errors.Is(err, gorm.ErrRecordNotFound) {
		roles := []*model.Role{
			{
				RoleID:   consts.RoleIDAdmin,
				RoleName: consts.RoleNameAdmin,
			},
			{
				RoleID:   consts.RoleIDTeacher,
				RoleName: consts.RoleNameTeacher,
			},
			{
				RoleID:   consts.RoleIDStudent,
				RoleName: consts.RoleNameStudent,
			},
		}
		for _, role := range roles {
			_ = role.Create()
		}
	}
	redis.InitRedis()
}

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
		return
	})

	user.Route(server)

	if err := server.Run(config.GetServer().Host + ":" + config.GetServer().Port); err != nil {
		panic(fmt.Errorf("服务运行失败 %s", err))
	}
}
