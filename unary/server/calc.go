package main

import (
	"context"
	"log"

	"github.com/cyberdr0id/go-grpc-practice/unary/proto"
)

func (s *Server) Sum(ctx context.Context, in *proto.UnaryRequest) (*proto.UnaryResponse, error) {
	log.Printf("Sum() function was invoked => %v\n", in)

	return &proto.UnaryResponse{
		Result: in.X + in.Y,
	}, nil
}
