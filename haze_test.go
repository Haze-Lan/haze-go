package main

import (
	"context"
	"github.com/Haze-Lan/haze-go/api"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)
const (
	address     = "localhost:50051"
	defaultName = "world"
)
func Test(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &api.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}

