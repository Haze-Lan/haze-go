package logger

import (
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var log_ *zapGrpcLoggerV2
var one sync.Once

type zapGrpcLoggerV2 struct {
	logger    *zap.Logger
	verbosity int
}

func (l *zapGrpcLoggerV2) Info(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}

func (l *zapGrpcLoggerV2) Infoln(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}

func (l *zapGrpcLoggerV2) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *zapGrpcLoggerV2) Warning(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}

func (l *zapGrpcLoggerV2) Warningln(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}

func (l *zapGrpcLoggerV2) Warningf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *zapGrpcLoggerV2) Error(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

func (l *zapGrpcLoggerV2) Errorln(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

func (l *zapGrpcLoggerV2) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

func (l *zapGrpcLoggerV2) Fatal(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}

func (l *zapGrpcLoggerV2) Fatalln(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}

func (l *zapGrpcLoggerV2) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(format, args...))
}

func (l *zapGrpcLoggerV2) V(level int) bool {
	return l.verbosity <= level
}
