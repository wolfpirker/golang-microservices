package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/wolfpirker/golang-microservices/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello I am a client")
	// to connect client, there is several options: we use grpc.WithInsecure()
	// since ssl certificates are a bit tricky to setup - it will be done later

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // certificate authority trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")

		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)

	// doServerStreaming(c)

	// doClientStreaming(c)

	// doBiDiStreaming(c)
	// doUnaryWithDeadline(c, 5*time.Second) // should complete
	// doUnaryWithDeadline(c, 1*time.Second) // should timeout
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a uanry RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Wolfgang",
			LastName:  "Pirker",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Wolfgang",
			LastName:  "Pirker",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
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
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Wolfgang",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Saven",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Dani",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Stefan",
			},
		},
	}

	// Note: we don't need a request
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	// we iterate over our slice and send each message
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(250 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)

}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Wolfgang",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Saven",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Dani",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Stefan",
			},
		},
	}

	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
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
		for _, req := range requests {
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
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc) // -> unblocks everything
	}()

	// block until everything is done
	<-waitc
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("starting to do a uanryWithDeadline RPC...")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Wolfgang",
			LastName:  "Pirker",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {

		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline exceeded!")
			} else {
				fmt.Println("unexpected error %v\n", statusErr)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		}
		return
	}
	log.Printf("Response from Greet: %v\n", res.Result)
}
