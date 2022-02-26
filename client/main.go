package main

import (
	"context"
	"flag"
	"fmt"

	sp "github.com/EvgenyiK/ClockwiseSpiral/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := sp.NewClockWiseSpiralServiceClient(conn)

	var a int
	flag.IntVar(&a, "a", 0, "The integer param")
	flag.Parse()

	resp, err := client.GenerateSpiral(context.Background(), &sp.NumbRequest{Num: int32(a)})
	if err != nil {
		logrus.Fatalf("could not get answer: %v", err)
	}

	for _, v := range resp.Responses {
		fmt.Println(v.Mid)
	}
}
