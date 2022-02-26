package main

import (
	"context"
	"fmt"
	"net"

	sp "github.com/EvgenyiK/ClockwiseSpiral/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ClockWiseSpiralServiceServer struct {
	sp.UnimplementedClockWiseSpiralServiceServer
}

func (c *ClockWiseSpiralServiceServer) GenerateSpiral(ctx context.Context, req *sp.NumbRequest) (*sp.NumbResponse, error) {
	var err error
	response := new(sp.NumbResponse)

	top, left, bottom, right := int32(0), int32(0), req.Num-1, req.Num-1
	if req.Num < 1 {
		return response, nil
	}

	size := req.Num * req.Num
	eq := make([]int32, size)

	i := int32(1)

	for left < right {

		for c := left; c <= right; c++ {
			eq[top*req.Num+c] = i
			i++
		}
		top++

		for r := top; r <= bottom; r++ {
			eq[r*req.Num+right] = i
			i++
		}
		right--

		if top == bottom {
			break
		}

		for c := right; c >= left; c-- {
			eq[bottom*req.Num+c] = i
			i++
		}
		bottom--

		for r := bottom; r >= top; r-- {
			eq[r*req.Num+left] = i
			i++
		}
		left++

	}

	eq[top*req.Num+left] = i

	for i := 0; i < len(eq)/int(req.Num)+1; i++ {
		f := int32(i) * req.Num
		l := f + req.Num
		if int(l) > len(eq) {
			l = int32(len(eq))
		}
		if f == l {
			break
		}

		mid := new(sp.MiddleRequest)

		mid.Mid = eq[f:l]

		response.Responses = append(response.Responses, mid)
	}

	return response, err
}

func main() {
	server := grpc.NewServer()
	instance := new(ClockWiseSpiralServiceServer)
	sp.RegisterClockWiseSpiralServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatal("Unable to create grpc listener:", err)
	}

	fmt.Println("start server on port 8080")

	if err = server.Serve(listener); err != nil {
		logrus.Fatal("Unable to start server:", err)
	}
}
