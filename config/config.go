package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config 包含了应用的所有配置
type AppConfig struct {
	RunMode  string `mapstructure:"RunMode"`
	HttpPort int    `mapstructure:"HttpPort"`
	// 添加其他需要的配置项
}

// Config变量将用于存储所有的配置信息
var Config AppConfig

// InitConfig 初始化配置
func InitConfig() {
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 设置配置文件所在的路径
	viper.AddConfigPath("./config")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// 将配置绑定到Config结构体
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}
