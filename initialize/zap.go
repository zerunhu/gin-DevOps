package initialize

import (
	"fmt"
	"gin-DevOps/config"
	"gin-DevOps/utils"
	rotateLogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var level zapcore.Level

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02-15:04:05"))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}


//日志文件切割，按天
func getWriter() (zapcore.WriteSyncer, error) {
	// 保存30天内的日志，每24小时(整点)分割一次日志
	fileWriter, err := rotateLogs.New(
		path.Join(config.GdoConfig.Zap.Director, "%Y-%m-%d.log"),
		rotateLogs.WithMaxAge(time.Hour*24*30),
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	if config.GdoConfig.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}


func getEncoderCore() (core zapcore.Core) {
	writer, err := getWriter() // 使用file-rotateLogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}


func Zap() *zap.Logger {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(config.GdoConfig.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", config.GdoConfig.Zap.Director)
		_ = os.Mkdir(config.GdoConfig.Zap.Director, os.ModePerm)
	}

	// 初始化配置文件的Level
	switch config.GdoConfig.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	logger := zap.New(getEncoderCore(), zap.AddCaller())
	return logger
	//return logger.Sugar()
}