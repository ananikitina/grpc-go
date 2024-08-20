package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ananikitina/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Calculator Client")

	cc, err := grpc.NewClient("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {

	fmt.Println("Starting to do Sum RPC")

	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC:%v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}
