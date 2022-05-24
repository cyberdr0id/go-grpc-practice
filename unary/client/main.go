package main

import (
	"fmt"
	"log"

	"github.com/cyberdr0id/go-grpc-practice/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:7777"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}
	defer conn.Close()

	c := proto.NewUnaryServiceClient(conn)

	res, err := doSum(c)
	if err != nil {
		log.Fatalf("Cannot find sum: %v\n", err)
	}

	fmt.Println("Result ", res)
}
