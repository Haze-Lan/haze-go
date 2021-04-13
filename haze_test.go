package main

import (
	"fmt"

	"google.golang.org/grpc"
	"log"
	"testing"
)

const (
	address = "localhost:80"
)

func Test(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", "", ""), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

}
