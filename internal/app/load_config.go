package app

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

// Config 对应 config.yml 的顶层结构。
type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

type Server struct {
	Host      string  `yaml:"host"`
	Port      int     `yaml:"port"`
	Name      string  `yaml:"name"`
	LogLevel  string  `yaml:"logLevel"`
	LogDir    string  `yaml:"logDir"`
	LogMaxAge int     `yaml:"logMaxAge"`
	Session   Session `yaml:"session"`
}

type Session struct {
	Sec      string `yaml:"sec"`
	Name     string `yaml:"name"`
	LastLang int    `yaml:"lastLang"`
}

type DB struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// Load 读取并解析 YAML 配置文件。
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}

	// 初始化默认配置
	var cfg Config = Config{
		Server: Server{
			Host:      "127.0.0.1",
			Port:      80,
			Name:      "app",
			LogLevel:  "info",
			LogDir:    "./logs",
			LogMaxAge: 7,
			Session: Session{
				Sec:      "default-sec",
				Name:     "session",
				LastLang: 1,
			},
		},
		DB: DB{
			Type:     "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Database: "app",
		},
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", path, err)
	}

	return &cfg, nil
}
