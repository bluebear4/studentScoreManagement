package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"studentScoreManagement/config"
)

var db *gorm.DB

func GetDatabase() *gorm.DB {
	return db
}

func InitDatabase() {
	//读取配置
	config := config.GetDatabase()

	//拼接dsn参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		config.Username, config.Password, config.Host, config.Port, config.Dbname, config.Timeout)

	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

}
