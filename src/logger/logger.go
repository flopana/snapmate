package logger

import (
	"fmt"
	"snapmate/config"
	"time"
)

type Logger struct {
	debugOutput bool
}

func newLogger(debugOutput bool) *Logger {
	return &Logger{debugOutput: debugOutput}
}

func NewLogger() *Logger {
	conf := config.GetConfig()
	return newLogger(conf.DebugLog)
}

func getTimestamp() string {
	return time.Now().Format(time.DateTime)
}

func (l *Logger) Debug(msg ...interface{}) {
	if !l.debugOutput {
		return
	}

	fullMsg := fmt.Sprint(msg...)
	m := fmt.Sprintf("[%s] [DEBUG] %s", getTimestamp(), fullMsg)
	fmt.Println(m)
}

func (l *Logger) Info(msg ...interface{}) {
	fullMsg := fmt.Sprint(msg...)
	m := fmt.Sprintf("[%s] [INFO] %s", getTimestamp(), fullMsg)
	fmt.Println(m)
}

func (l *Logger) Warn(msg ...interface{}) {
	fullMsg := fmt.Sprint(msg...)
	m := fmt.Sprintf("[%s] [WARN] %s", getTimestamp(), fullMsg)
	fmt.Println(m)
}

func (l *Logger) Error(msg ...interface{}) {
	fullMsg := fmt.Sprint(msg...)
	m := fmt.Sprintf("[%s] [ERROR] %s", getTimestamp(), fullMsg)
	fmt.Println(m)
}
