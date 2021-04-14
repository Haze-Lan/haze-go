package logger

import (
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
	"os"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewLogger() {
	path, _ := utils.GetCurrentDirectory()
	option.WithFilePath(path).Apply(option.LoggingOptionsInstance)
	/*hook := lumberjack.Logger{
		Filename:   opt.FilePath + "/app.log", //日志文件路径
		MaxSize:    128,                       //最大字节
		MaxAge:     30,
		MaxBackups: 7,
		Compress:   true,
	}*/
	w := zapcore.AddSync(os.Stdout)
	// debug->info->warn->error
	var level = zap.DebugLevel
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000Z0700")
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core)
	logger.Info("Logger init success")

	var log = &ZapLogger{
		logger: logger,
	}
	grpclog.SetLoggerV2(log)

}

func (zl *ZapLogger) Info(args ...interface{}) {
	zl.logger.Sugar().Info(args...)
}

func (zl *ZapLogger) Infoln(args ...interface{}) {
	zl.logger.Sugar().Info(args...)
}
func (zl *ZapLogger) Infof(format string, args ...interface{}) {
	zl.logger.Sugar().Infof(format, args...)
}

func (zl *ZapLogger) Warning(args ...interface{}) {
	zl.logger.Sugar().Warn(args...)
}

func (zl *ZapLogger) Warningln(args ...interface{}) {
	zl.logger.Sugar().Warn(args...)
}

func (zl *ZapLogger) Warningf(format string, args ...interface{}) {
	zl.logger.Sugar().Warnf(format, args...)
}

func (zl *ZapLogger) Error(args ...interface{}) {
	zl.logger.Sugar().Error(args...)
}

func (zl *ZapLogger) Errorln(args ...interface{}) {
	zl.logger.Sugar().Error(args...)
}

func (zl *ZapLogger) Errorf(format string, args ...interface{}) {
	zl.logger.Sugar().Errorf(format, args...)
}

func (zl *ZapLogger) Fatal(args ...interface{}) {
	zl.logger.Sugar().Fatal(args...)
}

func (zl *ZapLogger) Fatalln(args ...interface{}) {
	zl.logger.Sugar().Fatal(args...)
}

// Fatalf logs to fatal level
func (zl *ZapLogger) Fatalf(format string, args ...interface{}) {
	zl.logger.Sugar().Fatalf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (zl *ZapLogger) V(v int) bool {
	return false
}
