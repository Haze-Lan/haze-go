package main

import (
	"github.com/Haze-Lan/haze-go/examples/simple/haze-common/endpoint"
	"github.com/Haze-Lan/haze-go/examples/simple/haze-provider/impl"
	"github.com/Haze-Lan/haze-go/server"
	"log"
)

func main() {
	haze := server.NewServer()
	haze.RegisterService(endpoint.Account_ServiceDesc,&impl.AccountService{})

		if err := haze.Start(); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}




}