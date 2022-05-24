package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/cyberdr0id/go-grpc-practice/unary/proto"
)

func doSum(c proto.UnaryServiceClient) (int32, error) {
	log.Println("doSum() was invoked")

	rand.Seed(time.Now().UnixNano())

	res, err := c.Sum(context.Background(), &proto.UnaryRequest{
		X: int32(rand.Intn(20)) + 1,
		Y: int32(rand.Intn(20)) + 1,
	})
	if err != nil {
		return -1, fmt.Errorf("unable to find sum: %w", err)
	}

	return res.Result, nil
}
