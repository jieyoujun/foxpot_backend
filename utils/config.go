package utils

import (
	"gopkg.in/ini.v1"
)

// Config 配置项结构体
type Config struct {
	Foxpot  foxpotConfig  `ini:"foxpot"`
	DB      dbConfig      `ini:"db"`
	Session sessionConfig `ini:"session"`
}

type foxpotConfig struct {
	Address string `ini:"addr"`
	LogPath string `ini:"log_path"`
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

// GlobalConfig 全局配置
var GlobalConfig Config

// LoadConfigFile 从配置文件加载配置
func LoadConfigFile(filePath string) error {
	return ini.MapTo(&GlobalConfig, filePath)
}
