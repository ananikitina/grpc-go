package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ananikitina/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Hi, client")

	cc, err := grpc.NewClient("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do Unary RPC")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ana",
			LastName:  "Niki",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC:%v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
