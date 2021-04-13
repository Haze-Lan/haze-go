package server

import (
	"os"
	"os/signal"
	"syscall"
)

// Signal Handling
func (s *Server) handleSignals() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-c
		log.Info("Process exit instruction received")
		s.Shutdown()
	}()
}
