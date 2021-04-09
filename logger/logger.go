package logger

import (
	"github.com/Haze-Lan/haze-go/option"
	"google.golang.org/grpc/grpclog"
)

var loggingOptions *option.LoggingOptions

func init() {
	loggingOptions, _ = option.LoadLoggingOptions()
}

//Logger生成器
type Logger struct {
	name string
}

var cache = map[string]*Logger{}

func (c *Logger) Info(format string, args ...interface{}) {

	if len(args) == 0 {
		grpclog.Infoln(format)
	} else {
		grpclog.Infof(format, args)
	}
}

func (c *Logger) Warn(format string, args ...interface{}) {
	if len(args) == 0 {
		grpclog.Warningln(format)
	} else {
		grpclog.Warningf(format, args)
	}
}

func (c *Logger) Error(format string, args ...interface{}) {
	if len(args) == 0 {
		grpclog.Errorln(format)
	} else {
		grpclog.Errorf(format, args)
	}
}

func (c *Logger) Fatal(format string, args ...interface{}) {
	if len(args) == 0 {
		grpclog.Fatalln(format)
	} else {
		grpclog.Fatalf(format, args)
	}
}

func (c *Logger) V(l int) bool {
	return c.V(l)
}

func LoggerFactory(name string) *Logger {
	if cData, ok := cache[name]; ok {
		return cData
	}
	c := &Logger{name}
	cache[name] = c
	return c
}
