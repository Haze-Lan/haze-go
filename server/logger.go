package server
import (
	"fmt"
	"io"
	"os"
	"sync/atomic"

)


type Logger interface {

	// Log a notice statement
	Noticef(format string, v ...interface{})
	// Log a warning statement
	Warnf(format string, v ...interface{})
	// Log a fatal error
	Fatalf(format string, v ...interface{})
	// Log an error
	Errorf(format string, v ...interface{})
	// Log a debug statement
	Debugf(format string, v ...interface{})
	// Log a trace statement
	Tracef(format string, v ...interface{})
}



func (s *Server) ConfigureLogger() {
	s.loggers = make([]Logger,3)
    for  _,channel:=range s.opts.logger.channels {
		switch channel {
		case "CONSOLE":
			append(s.loggers, logger.NewConsoleLogger(s.opts.logger.level))
		case "FILE":

		case "REMOTE":

		}
	}
}



// Noticef logs a notice statement
func (s *Server) Noticef(format string, v ...interface{}) {
	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Noticef(format, v...)
	}, format, v...)
}

// Errorf logs an error
func (s *Server) Errorf(format string, v ...interface{}) {
	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Errorf(format, v...)
	}, format, v...)
}

// Error logs an error with a scope
func (s *Server) Errors(scope interface{}, e error) {
	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Errorf(format, v...)
	}, "%s - %s", scope, UnpackIfErrorCtx(e))
}

// Error logs an error with a context
func (s *Server) Errorc(ctx string, e error) {
	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Errorf(format, v...)
	}, "%s: %s", ctx, UnpackIfErrorCtx(e))
}

// Error logs an error with a scope and context
func (s *Server) Errorsc(scope interface{}, ctx string, e error) {
	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Errorf(format, v...)
	}, "%s - %s: %s", scope, ctx, UnpackIfErrorCtx(e))
}

// Warnf logs a warning error
func (s *Server) Warnf(format string, v ...interface{}) {
	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Warnf(format, v...)
	}, format, v...)
}

// Fatalf logs a fatal error
func (s *Server) Fatalf(format string, v ...interface{}) {
	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Fatalf(format, v...)
	}, format, v...)
}

// Debugf logs a debug statement
func (s *Server) Debugf(format string, v ...interface{}) {
	if atomic.LoadInt32(&s.logging.debug) == 0 {
		return
	}

	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Debugf(format, v...)
	}, format, v...)
}

// Tracef logs a trace statement
func (s *Server) Tracef(format string, v ...interface{}) {
	if atomic.LoadInt32(&s.opts.logger.trace) == 0 {
		return
	}
	s.executeLogCall(func(logger Logger, format string, v ...interface{}) {
		logger.Tracef(format, v...)
	}, format, v...)
}

func (s *Server) executeLogCall(f func(logger Logger, format string, v ...interface{}), format string, args ...interface{}) {
	if len(s.loggers) == 0 {
		fmt.Errorf("未初始化任何日志通道")
		return
	}
	for _,v:=range s.loggers{
		f(v, format, args...)
	}
}