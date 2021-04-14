package main

import (
	"context"
	"fmt"
	endpoint2 "github.com/Haze-Lan/haze-go/examples/haze-provider/endpoint"
	model2 "github.com/Haze-Lan/haze-go/examples/haze-provider/model"
	"github.com/Haze-Lan/haze-go/server"
	"log"
)

func main() {
	haze := server.NewServer()
	go func() {
		if err := haze.Start(); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	clien := endpoint2.NewAccountClient(haze.GetService("account-haze"))
	fmt.Println(clien.Authentication(context.TODO(), &model2.LoginRequest{Name: "22222", Pass: "22222222222222222"}))
}
