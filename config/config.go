package config

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

type Config struct {
	*Server
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var config Config

func GetConfig() Config {
	return config
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()

	if err != nil { // 处理错误
		panic(fmt.Errorf("配置文件初始化失败: %s \n", err))
	} else if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("配置文件序列化失败: %s \n", err))
	}

	//遍历结构体所有成员检查
	typeOfConfig := reflect.TypeOf(config)
	valueOfConfig := reflect.ValueOf(config)
	for i := 0; i < typeOfConfig.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfConfig.Field(i)
		fieldValue := valueOfConfig.Field(i)
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			panic(fmt.Errorf("未配置%s\n", fieldType.Name))
		}
	}
}
