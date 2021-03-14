package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wolfpirker/golang-microservices/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
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
}

func doUnary(c calcpb.CalcServiceClient) {

	var num1 int32 = 0
	var num2 int32 = 0
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
