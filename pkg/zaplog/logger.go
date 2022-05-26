package zaplog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"strings"
	"time"
)

func GetEncoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     getEncodeTime, // 自定义输出时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(config)
}

// 定义日志输出时间格式
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func GetLumberWriter(path string, maxSize, maxAge int, mode string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   getLogFile(path),
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: 5,
		LocalTime:  true,
		Compress:   false,
	}
	if mode == "PROD" {
		return zapcore.AddSync(lumberJackLogger)
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func GetLevel(level string) zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
		"fatal": zapcore.FatalLevel,
	}
	if l, ok := levelMap[level]; ok {
		return l
	}
	return zapcore.InfoLevel
}

func getLogFile(p string) string {
	timeFormat := time.Now().Format("2006-01-02")
	fileName := strings.Join([]string{
		timeFormat,
		"log",
	}, ".")
	return path.Join(p, fileName)
}
