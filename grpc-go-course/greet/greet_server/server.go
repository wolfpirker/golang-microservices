package main

import (
	"log"
	"net"
	"fmt"
	"context"

	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
)

type server struct{}

func main() {
	fmt.Println("Hello World")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
}
