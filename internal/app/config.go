package app

import (
	"log"
)

var cfg *Config

// 获取配置
func GetConfig() *Config {
	if cfg != nil {
		return cfg
	} else {
		c, err := Load("./config/config.yml")
		if err != nil {
			log.Fatal(err)
		}

		cfg = c
		return cfg
	}
}
