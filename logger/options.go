package logger

import "sync"

var defaultLoggerOptions = &loggerOptions{level: InfoLevel}

//日志配置
type loggerOptions struct {
	//日志输出通道 FILE CONSOLE REMOTE
	channels []string
	//文件日志分割周期
	limit int64
	//日志分割大小
	limitSize int64
	sync.Mutex
	//日志级别
	level Level
}

type LoggerOptions interface {
	apply(*loggerOptions)
}
type funcLoggerOptions struct {
	f func(*loggerOptions)
}

func (fdo *funcLoggerOptions) apply(do *loggerOptions) {
	fdo.f(do)
}

func newFuncServerOption(f func(*loggerOptions)) *funcLoggerOptions {
	return &funcLoggerOptions{
		f: f,
	}
}
