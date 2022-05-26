package main

import (
	"io"
	"log"

	"github.com/cyberdr0id/go-grpc-practice/cstream/proto"
)

func (s *Server) Avg(stream proto.AvgService_AvgServer) error {
	log.Println("Avg() was invoked")

	var res int32
	var i int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.AvgResponse{
				Result: float64(res) / float64(i),
			})
		}
		if err != nil {
			log.Fatalf("cannot recive data from client: %v\n", err)
		}

		log.Printf("Recive data => %d\n", req.Input)

		res += req.Input
		i++
	}
}
