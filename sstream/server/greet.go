package main

import (
	"log"

	"github.com/cyberdr0id/go-grpc-practice/sstream/proto"
)

func (s *Server) Prime(in *proto.PrimeRequest, stream proto.PrimeService_PrimeServer) error {
	log.Printf("Prime function was invoked => %v\n", in)

	k := int32(2)
	n := in.Input

	for n > 1 {
		if n%k == 0 {
			stream.Send(&proto.PrimeResponse{
				Result: k,
			})

			n /= k
		} else {
			k++
		}
	}

	return nil
}
