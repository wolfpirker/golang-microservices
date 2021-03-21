package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/wolfpirker/golang-microservices/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (*server) PrimeNumberDecomposition(in *calcpb.PrimeNumberDecompositionRequest, stream calcpb.CalcService_PrimeNumberDecompositionServer) error {
	// example: The client will send one number (120) and the server will respond
	// with a stream of (2,2,2,3,5), because 120=2*2*2*3*5

	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", in)

	k := int32(2)
	for N := in.GetNumber(); N > 1; {
		if mod(N, k) == 0 {
			res := &calcpb.PrimeNumberDecompositionResponse{
				Number: k,
			}
			N = N / k

			stream.Send(res)
		} else {
			k++
		}
	}

	return nil
}

func (*server) ComputeAverage(stream calcpb.CalcService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with a streaming request\n")

	sum := int32(0)
	for count := int32(0); ; count++ {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream

			average := float64(sum) / float64(count)
			return stream.SendAndClose(&calcpb.ComputeAverageResponse{
				Result: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		sum += req.GetNumber()
	}
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

func (*server) FindMaximum(stream calcpb.CalcService_FindMaximumServer) error {
	fmt.Printf("FindMaximum function was invoked with a streaming request\n")

	maximum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		num := req.GetNumber()
		if num > maximum {
			maximum = num

		}
		sendErr := stream.Send(&calcpb.FindMaximumResponse{
			Result: maximum,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}

// handson #44, error codes exercise
func (*server) SquareRoot(ctx context.Context, req *calcpb.SquareRootRequest) (*calcpb.SquareRootResponse, error) {
	fmt.Println("Received SquareRoot RPC")
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &calcpb.SquareRootResponse{
		Number: math.Sqrt(float64(number)),
	}, nil
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
