package app

import (
	"main/internal/utils/logger"
	"time"

	"go.uber.org/zap"
)

var appLog *zap.Logger

// var appLog *zap.Logger = zap.NewNop()

func Logger() *zap.Logger {
	if appLog != nil {
		return appLog
	}
	// 配置日志
	cfg := logger.DefaultConfig()
	cfg.Level = GetConfig().Server.LogLevel
	cfg.Directory = GetConfig().Server.LogDir
	cfg.Filename = GetConfig().Server.Name
	cfg.MaxAge = time.Duration(GetConfig().Server.LogMaxAge) * 24 * time.Hour

	l, err := logger.New(cfg)
	if err != nil {
		panic(err)
	}
	appLog = l
	defer appLog.Sync()
	return appLog
}
