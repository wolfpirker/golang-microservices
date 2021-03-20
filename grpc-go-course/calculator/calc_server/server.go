package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/wolfpirker/golang-microservices/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	firstNum := req.GetSummand1()
	secondNum := req.GetSummand2()
	res := &calcpb.SumResponse{
		SumResult: firstNum + secondNum,
	}
	return res, nil
}

// func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
func (*server) DecomposeNumber(in *calcpb.PrimeNumberDecompositionRequest, stream calcpb.CalcService_DecomposeNumberServer) error {
	// example: The client will send one number (120) and the server will respond
	// with a stream of (2,2,2,3,5), because 120=2*2*2*3*5

	var k int32 = 2
	for N := in.Number; N > 1; {
		if mod(N, k) == 0 {
			res := &calcpb.PrimeNumberDecompositionResponse{
				Number: k,
			}
			N = N / k

			stream.Send(res)
			time.Sleep(500 * time.Millisecond) // -> just for demonstration!
		} else {
			k++
		}
	}

	return nil
}

func mod(a, b int32) int32 {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func main() {
	fmt.Println("Calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
