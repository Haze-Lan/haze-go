package main

import (
	"github.com/Haze-Lan/haze-go/server"
	_ "github.com/Haze-Lan/haze-go/server"
	"log"
)

func main() {
	haze := server.NewServer()
	if err := haze.Start(); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
