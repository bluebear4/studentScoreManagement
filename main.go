package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"studentScoreManagement/config"
	"studentScoreManagement/db"
	"studentScoreManagement/handle"
	"studentScoreManagement/handle/user"
	"studentScoreManagement/model"
	"studentScoreManagement/redis"
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
		RoleID:   "0",
		RoleName: "admin",
	}
	if err := role.Find(); errors.Is(err, gorm.ErrRecordNotFound) {
		roles := []*model.Role{
			{
				RoleID:   "0",
				RoleName: "admin",
			},
			{
				RoleID:   "1",
				RoleName: "teacher",
			},
			{
				RoleID:   "2",
				RoleName: "student",
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
	server.GET("/ping", handle.Ping)

	User := server.Group("/user")
	{
		User.POST("/register", user.Register)
		User.POST("/login", user.Login)
	}

	if err := server.Run(config.GetServer().Host + ":" + config.GetServer().Port); err != nil {
		panic(fmt.Errorf("服务运行失败 %s", err))
	}
}
