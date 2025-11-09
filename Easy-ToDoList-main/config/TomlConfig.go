package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type TomlConfig struct {
	Server ServerConfig
	Sqlite SqliteConfig
	Log    LogConfig
	Jwt    JwtConfig
}

type ServerConfig struct {
	Port string
}

type SqliteConfig struct {
	Driver         string
	DataSourceName string
}

type LogConfig struct {
	Level string
	Path  string
}

type JwtConfig struct {
	SecretKey string
	Expire    int64
}

var c TomlConfig

func Init() {
	var err error
	// 设置文件名
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("toml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("fail to unmarshal config", err))
		return
	}
}
func Get() TomlConfig {
	return c
}
