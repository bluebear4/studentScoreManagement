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
	"studentScoreManagement/info"
	"studentScoreManagement/middleware"
	"studentScoreManagement/model"
	"studentScoreManagement/redis"
	"studentScoreManagement/score"
	"studentScoreManagement/user"
)

func init() {
	//初始化
	config.InitConfig()
	db.InitDatabase()
	redis.InitRedis()

	//数据库
	err := db.GetDatabase().AutoMigrate(
		&model.Role{},
		&model.Score{},
		&model.User{},
		&model.UserID2RoleID{},
		&model.UserInfo{},
	)
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
				RoleCode: nil,
			},
			{
				RoleID:   consts.RoleIDTeacher,
				RoleName: consts.RoleNameTeacher,
				RoleCode: nil,
			},
			{
				RoleID:   consts.RoleIDStudent,
				RoleName: consts.RoleNameStudent,
				RoleCode: nil,
			},
		}
		for _, role := range roles {
			_ = role.Create()
		}
	}

	//hard code 设定管理员
	admin := &model.User{
		ID:       "admin",
		PassWord: "123456",
	}
	if err := admin.Find(); errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.GetDatabase().Transaction(func(tx *gorm.DB) error {
			if err := admin.Create(); err != nil {
				panic(err)
			}
			//角色绑定
			uid2role := &model.UserID2RoleID{
				UserID: admin.ID,
				RoleID: consts.RoleIDAdmin,
			}
			return uid2role.Create()
		})
		if err != nil {
			panic(err)
		}
	}
}

func route(server *gin.Engine) {

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
		return
	})

	User := server.Group("/user")
	{
		User.POST("/register", user.Register)
		User.POST("/login", user.Login)
		User.POST("/changePassword", user.ChangePassword)
	}

	Student := server.Group("/student", middleware.Auth(map[int]bool{
		consts.RoleIDTeacher: true,
		consts.RoleIDStudent: true,
		consts.RoleIDAdmin:   true,
	}))
	{
		Student.POST("/logout", user.Logout)

		Student.POST("/getScoresByID", score.GetScoresByID)
		Student.POST("/getScoresByClass", score.GetScoresByClass)

		Student.GET("/getClasses", info.GetClasses)
		Student.GET("/getSubjects", score.GetSubjects)
	}

	Teacher := server.Group("/teacher", middleware.Auth(map[int]bool{
		consts.RoleIDTeacher: true,
		consts.RoleIDAdmin:   true,
	}))
	{
		Teacher.POST("/addInfo", info.AddInfo)
		Teacher.POST("/updateInfo", info.UpdateInfo)
		Teacher.POST("/deleteInfo", info.DeleteInfo)

		Teacher.POST("/addScore", score.AddScore)
		Teacher.POST("/updateScore", score.UpdateScore)
		Teacher.POST("/deleteScore", score.DeleteScore)

		upload := Teacher.Group("/upload")
		{
			upload.POST("/info", info.UploadInfo)
			upload.POST("/score", score.UploadScore)
		}
	}

	Admin := server.Group("/admin", middleware.Auth(map[int]bool{
		consts.RoleIDAdmin: true,
	}))
	{
		Admin.POST("/changeValidateCode", user.ChangeValidateCode)
	}
}
func main() {
	server := gin.Default()

	route(server)

	if err := server.Run(config.GetServer().Host + ":" + config.GetServer().Port); err != nil {
		panic(fmt.Errorf("服务运行失败 %s", err))
	}
}
