package main

import (
	"log"
	"net"

	"github.com/cyberdr0id/go-grpc-ptactice/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %s\n", err)
	}

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s/n", err)
	}
}
