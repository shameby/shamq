package logger

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	. "shamq/config"
)

//Logger logger实例type
type Logger struct {
	err  *log.Entry
	info *log.Entry
}

var logger *Logger

// InitLogger 初始化Logger
func init() {
	loggerPath := Conf.Common.LogPATH + "shamq_log"

	_, err := os.Stat(loggerPath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(loggerPath, os.ModePerm); err != nil {
			panic(err)
		}
	}

	loggerInfo := log.New()
	loggerErr := log.New()

	loggerInfo.Out = openOrCreateF(loggerPath + "/" + "shamq_" + time.Now().Format("20060102") + ".log")
	loggerErr.Out = openOrCreateF(loggerPath + "/" + "shamq_error_" + time.Now().Format("20060102") + ".log")

	loggerInfo.Formatter = &log.JSONFormatter{}
	loggerErr.Formatter = &log.JSONFormatter{}

	logger = &Logger{
		log.NewEntry(loggerErr),
		log.NewEntry(loggerInfo),
	}
}

// openOrCreateF 如果文件不存在则创建，存在则打开
func openOrCreateF(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(path)
		if err != nil {
			panic(err)
		}
	}
	return f
}

// Info 打印Info类型日志
func Info(message string, args ...interface{}) {
	logger.info.Printf(message, args...)
}

// Error 打印error类型日志
func Error(message string, args ...interface{}) {
	logger.err.Errorf(message, args...)
}

// Panic 打印panic类型日志
func Panic(message string, args ...interface{}) {
	logger.err.Panicf(message, args...)
}

// Fatal 打印fatal类型日志
func Fatal(args ...interface{}) {
	logger.err.Fatal(args...)
}
