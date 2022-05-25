package main

import (
	"log"
	"os"
	"strconv"

	"github.com/cyberdr0id/go-grpc-practice/sstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Parameters number must be equal to 2")
	}

	address := "localhost:" + os.Args[1]

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}
	defer conn.Close()

	c := proto.NewPrimeServiceClient(conn)

	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("cannot conver input parameter: %v\n", err)
	}

	err = doPrime(c, int32(n))
	if err != nil {
		log.Fatalf("Error while reading stream: %v\n", err)
	}

}
