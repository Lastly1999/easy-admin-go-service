package viper

import (
	viper2 "github.com/spf13/viper"
	"os"
)

var SrvConf *ServiceConfig

func init() {
	var serviceConfig ServiceConfig
	conf := serviceConfig.getConf()
	SrvConf = conf
}

// ServiceConfig 项目配置结构体
type ServiceConfig struct {
	Port     int            `yaml:"port"`
	Database DataBaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}

// DataBaseConfig 数据库配置的结构体
type DataBaseConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}

// RedisConfig Redis的配置
type RedisConfig struct {
	Host     string `yaml:"host"`
	PassWord string `yaml:"password"`
	Db       int    `yaml:"db"`
}

/**
  getConf 读取Yaml配置文件，并转换成ServiceConfig对象  struct结构
*/
func (serviceConfig *ServiceConfig) getConf() *ServiceConfig {
	// 获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 初始化viper
	viper := viper2.New()
	// 添加要读取的文件路径
	viper.AddConfigPath(path + "/config")
	// 设置要读取的配置文件名称
	viper.SetConfigName("config.yaml")
	// 设置要读取的配置文件类型
	viper.SetConfigType("yaml")
	// 进行读取
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析配置文件 传入结构体
	if err = viper.Unmarshal(&serviceConfig); err != nil {
		// 解析配置文件失败
		panic(err)
	}
	// 读取成功
	return serviceConfig
}
