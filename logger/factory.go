package logger

import (
	"log"
	"os"
)

var cache = make(map[string]*NameLogger)

func Factory(name string) *NameLogger {
	if t, ok := cache[name]; ok {
		return t
	}
	var log_ = log.New(os.Stdout, "["+name+"] ", log.LstdFlags)
	var i = &NameLogger{name: name, log: log_}
	cache[name] = i
	return i
}

type NameLogger struct {
	name string
	log  *log.Logger
}

func (s *NameLogger) Info(format string, args ...interface{}) {
	s.log.Printf(format, args...)
}
func (s *NameLogger) Warning(format string, args ...interface{}) {
	s.log.Printf(format, args...)
}
func (s *NameLogger) Error(format string, args ...interface{}) {
	s.log.Printf(format, args...)
}
func (s *NameLogger) Fatal(format string, args ...interface{}) {
	s.log.Fatalf(format, args...)
}
