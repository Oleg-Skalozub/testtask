package logger

import (
	"log"
	"runtime"
	"strconv"
)

const ERROR = " ERROR "
const DEBUG = " DEBUG "

// Logger ...
type Logger interface {
	Error(v ...interface{})
	Debug(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	logMain *log.Logger
}

// Error ...
func (lg *logger) Error(v ...interface{}) {
	path := getCaller()
	lg.logMain.Println(path, ERROR, v)
}

// Debug ...
func (lg *logger) Debug(v ...interface{}) {
	path := getCaller()
	lg.logMain.Println(path, DEBUG, v)
}

// Println ...
func (lg *logger) Println(v ...interface{}) {
	lg.logMain.Println(v)
}

// Printf ...
func (lg *logger) Printf(format string, v ...interface{}) {
	lg.logMain.Println(v)
}

func getCaller() string {
	buf := make([]byte, 0)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "????"
		line = 0
	}

	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}

	buf = append(buf, short...)
	buf = append(buf, ':')
	buf = append(buf, []byte(strconv.Itoa(line))...)
	buf = append(buf, ':')

	return string(buf)
}
