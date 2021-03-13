package main

import (
	"fmt"
	"log"

	"github.com/wolfpirker/golang-microservices/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")
	// to connect client, there is several options: we use grpc.WithInsecure()
	// since ssl certificates are a bit tricky to setup - it will be done later
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
