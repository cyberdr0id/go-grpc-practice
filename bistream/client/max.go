package main

import (
	"context"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/cyberdr0id/go-grpc-practice/bistream/proto"
)

func sliceAtoi(sa []string) ([]int32, error) {
	si := make([]int32, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, int32(i))
	}
	return si, nil
}

func doMax(c proto.MaxServiceClient, values []string) {
	v, err := sliceAtoi(values)
	if err != nil {
		log.Printf("cannot convert input data: %v\n", err)
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Printf("cannot create stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, n := range v {
			stream.Send(&proto.MaxRequest{
				Input: n,
			})

			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error while receiving: %v\n", err)
			}

			log.Printf("Current max value => %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
