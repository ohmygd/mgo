package log

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
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
			logPath: "./mLog",
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
	l.mu.Lock()
	defer l.mu.Unlock()

	_, fileName, line, ok := runtime.Caller(0)

	if !ok {
		//todo error 返回
		fileName, line = "", 0
	}

	info := []string{
		time.Now().Format("2006-01-02 15:04:05"),
		string(level),
		fileName,
		strconv.Itoa(line),
		format,
	}
	content := strings.Join(info, " ")

	_, _ = fmt.Fprintf(l.file, content, args...)
}

func (l *Logger)Close() {
	// todo 暂无用
	l.file.Close()
}

