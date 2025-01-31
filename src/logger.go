package main

import (
	"fmt"
	"time"
)

type logger struct {
	debugOutput bool
}

func newLogger(debugOutput bool) *logger {
	return &logger{debugOutput: debugOutput}
}

func getTimestamp() string {
	return time.Now().Format(time.DateTime)
}

func (l *logger) debug(msg ...interface{}) {
	if !l.debugOutput {
		return
	}

	fullMsg := fmt.Sprint(msg...)
	m := fmt.Sprintf("[%s] [DEBUG] %s", getTimestamp(), fullMsg)
	fmt.Println(m)
}

func (l *logger) info(msg ...interface{}) {
	fullMsg := fmt.Sprint(msg...)
	m := fmt.Sprintf("[%s] [INFO] %s", getTimestamp(), fullMsg)
	fmt.Println(m)
}

func (l *logger) warn(msg ...interface{}) {
	fullMsg := fmt.Sprint(msg...)
	m := fmt.Sprintf("[%s] [WARN] %s", getTimestamp(), fullMsg)
	fmt.Println(m)
}

func (l *logger) error(msg ...interface{}) {
	fullMsg := fmt.Sprint(msg...)
	m := fmt.Sprintf("[%s] [ERROR] %s", getTimestamp(), fullMsg)
	fmt.Println(m)
}
