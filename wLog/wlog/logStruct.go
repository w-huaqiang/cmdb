package wlog

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//LogLevel 定义日志级别类型
type LogLevel uint

// 定义日志级别
// UNKNOW 为0 : 未知日志级别
// Debug 为1 : debug日志级别，调试时使用
// Info 为2 : message 日志级别
// Warning 为3 : 警告级别，程序尚可运行，但是不是推荐方式
// Error 为4 : 错误级别，程序主体可运行，但是出现错误
// Fatal 为5 : 严重错误级别，程序不可运行
const (
	UNKNOW LogLevel = iota
	Debug
	Info
	Warning
	Error
	Fatal
)

// Wlog 定义日志输出结构体
// 包含日志输出的os.file 类型，和日志级别
type Wlog struct {
	Logfile *os.File
	Level   LogLevel
}

// 解析字符串为日志级别
func parseloglevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warning":
		return Warning
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return UNKNOW

	}
}

// 解析日志级别到字符串
func loglevelstring(level LogLevel) string {
	switch level {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return "UNKNOW"
	}
}

func loginfo(n int) (file string, object string, line int) {
	pc, file, line, ok := runtime.Caller(n)
	if ok {
		funcName := runtime.FuncForPC(pc)
		return file, funcName.Name(), line
	}

	return "null", "null", 0

}

// NewWlog 初始化一个日志记录
// 传入参数为 日志输出方式，记录等级；当日志超过或等于该等级才会被记录
func NewWlog(logFile interface{}, level string) Wlog {

	// 解析初始化日志的时候日志等级
	loglevel := parseloglevel(level)

	// 判断日志是写入文件还是一个可以直接使用的os.file
	switch logFile.(type) {
	case string:
		fileObj, err := os.OpenFile(logFile.(string), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("can't write log to %s", logFile.(string))
		}
		return Wlog{fileObj, loglevel}
	case *os.File:
		return Wlog{logFile.(*os.File), loglevel}
	default:
		return Wlog{}
	}

}

// Log wlog的记录方法，可以传入字符串，或者格式化
func (w *Wlog) Log(level LogLevel, format string, a ...interface{}) {
	// fmt.Println(level)
	// fmt.Println(w.Level)
	logstring := fmt.Sprintf(format, a...)
	file, object, line := loginfo(3)
	if level >= w.Level {
		logtime := time.Now()
		logObj := "[" + file + ":" + object + ":" + strconv.Itoa(line) + "]"
		logObj = logObj + "[" + loglevelstring(level) + "][" + logtime.Format("2006-01-02 15:04:05") + "] " + logstring + string('\n')
		w.Logfile.WriteString(logObj)
	}

}

// Debug 记录调试等级
func (w *Wlog) Debug(s string, a ...interface{}) {
	w.Log(Debug, s, a...)
}

// Info 记录基本信息等级
func (w *Wlog) Info(s string, a ...interface{}) {
	w.Log(Info, s, a...)
}

// Warning 记录告警等级
func (w *Wlog) Warning(s string, a ...interface{}) {
	w.Log(Warning, s, a...)
}

// Error 记录错误等级
func (w *Wlog) Error(s string, a ...interface{}) {
	w.Log(Error, s, a...)
}

// Fatal 记录严重错误等级
func (w *Wlog) Fatal(s string, a ...interface{}) {
	w.Log(Fatal, s, a...)
}

// Close 关闭日志记录
func (w *Wlog) Close() {
	w.Logfile.Close()
}
