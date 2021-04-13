// +build !windows

package server

import (
	"os"
	"os/signal"
	"syscall"
)

// Signal Handling
func (s *Server) handleSignals() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP)
	go func() {
		<-c
		log.Info("Process exit instruction received")
		s.Shutdown()
		os.Exit(1)
	}()
}
