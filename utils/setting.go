package utils

import (
	"gopkg.in/ini.v1"
)

// TODO
// 得有个检查空配置项给个默认值的东西
// 防止有非法的配置进来瞎搞

// Configuration 系统配置项
type Configuration struct {
	Foxpot  foxpotConfig  `ini:"foxpot"`
	DB      dbConfig      `ini:"db"`
	Session sessionConfig `ini:"session"`
}

type foxpotConfig struct {
	Address  string `ini:"addr"`
	LogPath  string `ini:"log_path"`
	LogLevel string `ini:"log_level"`
}

type dbConfig struct {
	Type        string `ini:"type"`
	DSN         string `ini:"dsn"`
	MaxOpenConn int    `ini:"max_open_conn"`
	MaxIdleConn int    `ini:"max_idle_conn"`
	MaxLifeTime int    `ini:"max_life_time"`
}

type sessionConfig struct {
	Key    string `ini:"key"`
	Secret string `ini:"secret"`
}

// Config 系统配置
var Config Configuration

// LoadConfigFile 加载配置文件
func LoadConfigFile(filePath string) error {
	return ini.MapTo(&Config, filePath)
}
