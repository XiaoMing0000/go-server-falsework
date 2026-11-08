package logger

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log 是全局 logger，New 成功后指向同一实例；未初始化时为 Nop，可直接调用。
var Log = zap.NewNop()

type Config struct {
	Level     string
	Directory string
	Filename  string
	MaxAge    time.Duration
}

func DefaultConfig() Config {
	return Config{
		Level:     "info",
		Directory: "/data/logs",
		Filename:  "app",
		MaxAge:    30 * 24 * time.Hour,
	}
}

func New(cfg Config) (*zap.Logger, error) {
	var level zapcore.Level

	if err := level.UnmarshalText([]byte(strings.ToLower(cfg.Level))); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(cfg.Directory, 0755); err != nil {
		return nil, err
	}

	pattern := filepath.Join(
		cfg.Directory,
		cfg.Filename+"-%Y-%m-%d.log",
	)

	writer, err := rotatelogs.New(
		pattern,
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(cfg.MaxAge),
	)
	if err != nil {
		return nil, err
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		MessageKey: "msg",

		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}

	// 终端和文件使用相同的 JSON 编码，保证输出格式一致。
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(writer),
		level,
	)
	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		level,
	)

	// 同时写入终端和文件。
	core := zapcore.NewTee(fileCore, consoleCore)

	l := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	Log = l

	return l, nil
}
