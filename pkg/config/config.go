package config

import (
	"goblog/pkg/logger"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

//Viper库 实例ring
var Viper *viper.Viper

type StrMap map[string]interface{}

func init() {

	//初始化
	Viper = viper.New()
	//设置文件名称
	Viper.SetConfigName(".env")
	// 3. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	Viper.SetConfigType("env")
	//环境变量配置文件查找的路径，相对于 main.go
	Viper.AddConfigPath(".")

	//开始读取 根目录下的 .env文件
	err := Viper.ReadInConfig()
	logger.LogError(err)

	//设置环境变量前缀，用以区分 Go 的系统环境变量
	Viper.SetEnvPrefix("appenv")
	//Viper.Get() 时，优先读取环境变量
	Viper.AutomaticEnv()

}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}

	return Get(envName)
}

//Add 新增配置项
func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

//Get 获取配置项，允许使用点式获取，如：app.name
func Get(path string, defaultValue ...interface{}) interface{} {
	//不存在的情况
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return Viper.Get(path)
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
