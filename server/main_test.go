package main

import (
	"context"
	"net"
	"os"
	"testing"

	sp "github.com/EvgenyiK/ClockwiseSpiral/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":50051"
)

func Server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sp.RegisterClockWiseSpiralServiceServer(s, &ClockWiseSpiralServiceServer{})
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func TestMain(m *testing.M) {
	go Server()
	os.Exit(m.Run())
}

func Equal(a, b []int32) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func TestNums(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sp.NewClockWiseSpiralServiceClient(conn)

	t.Run("Spiral", func(t *testing.T) {
		numb := int32(3)
		r, err := c.GenerateSpiral(context.Background(), &sp.NumbRequest{Num: numb})
		if err != nil {
			t.Fatalf("could not greet: %v", err)
		}

		var ex []*sp.MiddleRequest
		ex = append(ex, &sp.MiddleRequest{Mid: []int32{1, 2, 3}},
			            &sp.MiddleRequest{Mid: []int32{8, 9, 4}},
			            &sp.MiddleRequest{Mid: []int32{7, 6, 5}})

		for a, v := range r.Responses {
			for b, e := range ex {
				if a == b {
					j:= Equal(v.Mid,e.Mid)
					if j!= true {
						t.Error("Expected true, got ", j)
					}
				}
			}
		}

	})
}
