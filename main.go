package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentScoreManagement/config"
	"studentScoreManagement/handle"
)

func main() {
	fmt.Printf("配置读取结果为: %q\n", config.GetConfig())

	server := gin.Default()
	server.GET("/ping", handle.Ping)

	User := server.Group("/user")
	{
		User.GET("/login", handle.Login)
	}

	if err := server.Run(config.GetConfig().Host + ":" + config.GetConfig().Port); err != nil {
		panic(fmt.Errorf("服务运行失败 %s", err))
	}
}
