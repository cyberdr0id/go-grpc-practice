package main

import (
	"log"
	"net"
	"os"

	"github.com/cyberdr0id/go-grpc-practice/cstream/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.AvgServiceServer
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Parameters number must be equal to 2")
	}

	address := "0.0.0.0:" + os.Args[1]
	network := "tcp"

	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("Failed to listen server on %s: %s\n", address, err)
	}

	log.Printf("Listening on %s\n", address)

	s := grpc.NewServer()
	proto.RegisterAvgServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s\n", err)
	}
}
