package logging

import (
	"path"
	"time"

	"github.com/0x00-ketsu/lazypip/internal/config"
	"github.com/juju/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var conf *config.Config

func init() {
	conf, _ = config.Load()
}

func GetLogger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customEncodeTime
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	var filename string
	switch conf.Log.Level {
	case "debug":
		filename = "debug.log"
	case "info":
		filename = "info.log"
	case "error":
		filename = "error.log"
	default:
		filename = "info.log"
	}

	filepath := path.Join(conf.Log.Directory, filename)
	priority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.DebugLevel
	})

	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    50, // unit: MB
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   false,
	})

	var writeSyncer zapcore.WriteSyncer
	writeSyncer = zapcore.NewMultiWriteSyncer(fileWriteSyncer)
	fileCore := zapcore.NewCore(encoder, writeSyncer, priority)

	return zap.New(zapcore.NewTee(fileCore), zap.AddCaller())
}

// Adds custom content to EncodeTime
func customEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(conf.Log.Prefix + "2006-01-02 15:04:05.000"))
}
