package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cyberdr0id/go-grpc-practice/cstream/proto"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func doAvg(c proto.AvgServiceClient, values []string) error {
	v, err := sliceAtoi(values)
	if err != nil {
		return fmt.Errorf("cannot convert input number: %w", err)
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		return fmt.Errorf("cannot call Avg(): %w", err)
	}

	for _, n := range v {
		stream.Send(&proto.AvgRequest{
			Input: int32(n),
		})

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("cannot recive response from Avg(): %w", err)
	}

	log.Printf("Result => %.3f\n", res.Result)

	return nil
}
