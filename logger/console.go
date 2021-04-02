

package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type ConsoleLogger struct {
	sync.Mutex
	logger     *log.Logger
	debug      bool
	trace      bool
	infoLabel  string
	warnLabel  string
	errorLabel string
	fatalLabel string
	debugLabel string
	traceLabel string
	fl         *fileLogger
}

func NewConsoleLogger(level Level) ConsoleLogger {
	flags := 0
	flags = log.LstdFlags | log.Lmicroseconds
	pre := pidPrefix()
	l := &ConsoleLogger{
		logger: log.New(os.Stderr, pre, flags),
		debug:  true,
		trace:  true,
	}
	setColoredLabelFormats(l)
	return l
}


func setColoredLabelFormats(l *ConsoleLogger) {
	colorFormat := "[\x1b[%sm%s\x1b[0m] "
	l.infoLabel = fmt.Sprintf(colorFormat, "32", "INF")
	l.debugLabel = fmt.Sprintf(colorFormat, "36", "DBG")
	l.warnLabel = fmt.Sprintf(colorFormat, "0;93", "WRN")
	l.errorLabel = fmt.Sprintf(colorFormat, "31", "ERR")
	l.fatalLabel = fmt.Sprintf(colorFormat, "31", "FTL")
	l.traceLabel = fmt.Sprintf(colorFormat, "33", "TRC")
}

// Log a notice statement
func (log *ConsoleLogger) Noticef(format string, v ...interface{}){

}

// Log a warning statement
func (log *ConsoleLogger) Warnf(format string, v ...interface{}){

}

// Log a fatal error
func (log *ConsoleLogger) Fatalf(format string, v ...interface{}){

}

// Log an error
func (log *ConsoleLogger) Errorf(format string, v ...interface{}){

}

// Log a debug statement
func (log *ConsoleLogger) Debugf(format string, v ...interface{}){

}

// Log a trace statement
func (log *ConsoleLogger) Tracef(format string, v ...interface{}){

}