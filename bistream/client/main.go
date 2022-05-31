package main

import (
	"log"
	"os"

	"github.com/cyberdr0id/go-grpc-practice/bistream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("Parameters number must be more than 2")
	}

	address := "localhost:" + os.Args[1]

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}
	defer conn.Close()

	c := proto.NewMaxServiceClient(conn)

	doMax(c, os.Args[2:])
}
