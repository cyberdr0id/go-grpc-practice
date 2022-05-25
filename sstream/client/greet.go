package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/cyberdr0id/go-grpc-practice/sstream/proto"
)

func doPrime(c proto.PrimeServiceClient, num int32) error {
	log.Println("doGreet() was invoked")

	stream, err := c.Prime(context.Background(), &proto.PrimeRequest{
		Input: num,
	})
	if err != nil {
		return fmt.Errorf("cannot call Greet() function: %w", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error while reading stream: %w", err)
		}

		log.Printf("message from Prime() => %s\n", msg)

		time.Sleep(1 * time.Second)
	}

	return nil
}
