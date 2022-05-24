package main

import (
	"log"
	"net"
	"os"


	"github.com/cyberdr0id/go-grpc-practice/unary/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnaryServiceServer
}


var (
	address = "0.0.0.0:" + os.Args[1]
	network = "tcp"
)

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("Failed to listen server on %s: %s\n", address, err)
	}

	log.Printf("Listening on %s\n", address)

	s := grpc.NewServer()
	proto.RegisterUnaryServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s\n", err)
	}
}
