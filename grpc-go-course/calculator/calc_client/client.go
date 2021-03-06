package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/wolfpirker/golang-microservices/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("client")
	// to connect client, there is several options: we use grpc.WithInsecure()
	// since ssl certificates are a bit tricky to setup - it will be done later
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)
	doUnary(c)

	// doServerStreaming(c)

	// doClientStreaming(c)

	// doBiDiStreaming(c)

	doErrorUnary(c)
}

func doUnary(c calcpb.CalcServiceClient) {

	var num1 int32 = 0
	var num2 int32 = 0
	fmt.Println("#Unary streaming exercise")
	fmt.Println("enter a number 1 and 2 to let the server create a sum")
	_, err := fmt.Scanf("%d", &num1)
	_, err2 := fmt.Scanf("%d", &num2)
	if err != nil || err2 != nil {
		log.Fatalf("input cannot be parsed as integer")
		return
	}
	req := &calcpb.SumRequest{
		Summand1: num1,
		Summand2: num2,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calc RPC: %v", err)
	}
	log.Printf("Response from Calc Server: %v", res.SumResult)
}

func doServerStreaming(c calcpb.CalcServiceClient) {
	var num1 int32 = 0
	fmt.Println("#Server streaming exercise, DecomposeNumber")
	fmt.Println("enter a number: ")
	_, err := fmt.Scanf("%d", &num1)

	req := &calcpb.PrimeNumberDecompositionRequest{
		Number: num1,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling DecomposeNumber RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from DecomposeNumber: %v", msg.GetNumber())
	}

}

func doClientStreaming(c calcpb.CalcServiceClient) {
	var num1 int32 = 0
	fmt.Println("Starting to do a Client Streaming RPC...")

	// Note: we don't need a request
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling ComputeAverage: %v", err)
	}

	// we iterate over our slice and send each message
	for {
		fmt.Println("enter a number (or anything else to stop): ")
		_, err := fmt.Scanf("%d", &num1)

		if err != nil {
			// stop
			break
		} else {
			req := &calcpb.ComputeAverageRequest{
				Number: num1,
			}
			fmt.Printf("Sending req: %v\n", req)
			stream.Send(req)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from ComputeAverage: %v", err)
	}
	fmt.Printf("ComputeAverage Response: %v\n", res.GetResult())

}

func doBiDiStreaming(c calcpb.CalcServiceClient) {
	var num1 int32 = 0
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	// wait channel -> trick to block
	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		// Note: being run in its own goroutine! -> wouldn't have to
		// just to show things can really run in parallel
		for {
			fmt.Println("enter a number (or anything else to stop): ")
			_, err := fmt.Scanf("%d", &num1)

			if err != nil {
				// stop
				fmt.Println("->stop")
				break
			}
			req := &calcpb.FindMaximumRequest{
				Number: num1,
			}
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// we receive bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			fmt.Printf("Received Maximum: %v\n", res.GetResult())
		}
		close(waitc) // -> unblocks everything
	}()

	// block until everything is done
	<-waitc
}

func doErrorUnary(c calcpb.CalcServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")

	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -1)
}

func doErrorCall(c calcpb.CalcServiceClient, n int32) {

	res, err := c.SquareRoot(
		context.Background(),
		&calcpb.SquareRootRequest{
			Number: n,
		},
	)

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
		}
	}
	fmt.Printf("Result of square root of %v: %v\n", n, res.GetNumber())

}
