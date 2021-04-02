package main

import (
	"github.com/Haze-Lan/haze-go/server"
	_ "github.com/Haze-Lan/haze-go/server"
	"log"
)

func main() {
	haze := server.Init()
	if err := server.Run(haze); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
