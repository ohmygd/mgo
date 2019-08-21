package log

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var logger *Logger

type Level string

const (
	DebugLevel Level = "DEBUG"
	InfoLevel  = "Info"
	WarnLevel  = "Warn"
	ErrorLevel  = "Error"
)

type Logger struct {
	mu sync.Mutex
	logPath string
	logName string
	file *os.File
}

func init() {
	logger = &Logger{
			logPath: "/Users/machao/goBase/src/log",
			logName: "log",
		}

	filename := fmt.Sprintf("%s/%s.log", logger.logPath, logger.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755) // os.O_CREATE 创建文件 os.O_APPEND 追加写入 os.O_WRONLY 只写操作
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err: %v", filename, err))
	}

	logger.file = file
}

func Debug(format string, args ...interface{}) {
	Logf(DebugLevel, format, args...)
}

func Info(format string, args ...interface{}) {
	Logf(InfoLevel, format, args...)
}

func Warn(format string, args ...interface{}) {
	Logf(WarnLevel, format, args...)
}

func Error(format string, args ...interface{}) {
	Logf(ErrorLevel, format, args...)
}

func Logf(level Level, format string, args ...interface{}) {
	logger.Log(level, format, args...)
}

func (l *Logger)Log(level Level, format string, args ...interface{}) {
	info := []string{
		"[" + string(level) + "]",
		format,
	}
	content := strings.Join(info, " ") + "\n"

	dbgLogger := log.New(l.file, "", log.Llongfile | log.LstdFlags)

	//dbgLogger.Println(content)

	dbgLogger.Output(5, content)

}

func (l *Logger)Close() {
	// todo 暂无用
	l.file.Close()
}

