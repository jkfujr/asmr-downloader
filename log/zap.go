package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var AsmrLog *zap.Logger
var LogFile *os.File

const logDir = "./logs"

func init() {
	// 创建一个控制台的 encoder
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// 创建一个文件的 encoder
	fileEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	// 设置控制台的输出
	consoleOutput := zapcore.Lock(os.Stdout)

	_, err := os.Stat(logDir) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {

		} else {
			os.Mkdir(logDir, 0755)
		}
	}
	//当前时间
	now := time.Now()
	// Format the time using the standard format string
	currentTimeStr := now.Format("2006-01-02 15:04:05")

	// 设置日志文件的输出
	file, _ := os.OpenFile(logDir+string(filepath.Separator)+currentTimeStr+"asmr.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	//defer file.Close()
	fileOutput := zapcore.AddSync(file)

	// 设置日志级别和输出方式
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleOutput, zap.DebugLevel),
		zapcore.NewCore(fileEncoder, fileOutput, zap.DebugLevel),
	)

	// 创建 logger
	logger := zap.New(core)

	// 输出日志信息
	//logger.Info("hello world")
	AsmrLog = logger
	LogFile = file
}

func TestZapLog() {
	AsmrLog.Info("name", zap.String("info", "message"))
}
