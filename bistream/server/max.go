package main

import (
	"fmt"
	"io"
	"log"

	"github.com/cyberdr0id/go-grpc-practice/bistream/proto"
)

func (s *Server) Max(stream proto.MaxService_MaxServer) error {
	log.Println("Max() was invoked")

	var max int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("cannot recive data: %w", err)
		}

		if req.Input > max {
			max = req.Input
			err = stream.Send(&proto.MaxResponse{
				Result: max,
			})
			if err != nil {
				return fmt.Errorf("cannot send data: %w", err)
			}
		}

		log.Printf("Received: %d\n", req.Input)
	}
}
