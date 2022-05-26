package main

import (
	"log"
	"os"
	"strconv"

	"github.com/cyberdr0id/go-grpc-practice/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:" + os.Args[1]
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Parameters number must be equal to 3")
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}
	defer conn.Close()

	c := proto.NewUnaryServiceClient(conn)

	x, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid input parameter: %v\n", err)
	}

	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("Invalid input parameter: %v\n", err)
	}

	res, err := doSum(c, int32(x), int32(y))
	if err != nil {
		log.Fatalf("Cannot find sum: %v\n", err)
	}

	log.Println("Result ", res)
}
