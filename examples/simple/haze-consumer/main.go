package main

import (
	"context"
	"fmt"
	"github.com/Haze-Lan/haze-go/examples/simple/haze-common/endpoint"
	"github.com/Haze-Lan/haze-go/examples/simple/haze-common/model"
	"github.com/Haze-Lan/haze-go/server"
	"log"
)

func main() {
	haze := server.NewServer()
	clien := endpoint.NewAccountClient(haze.GetService("account-haze"))
	fmt.Println(clien.Authentication(context.TODO(), &model.LoginRequest{Name: "22222", Pass: "22222222222222222"}))

		if err := haze.Start(); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}



}
