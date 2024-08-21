package main

import (
	"context"
	"fmt"
	"io"
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

	//doUnary(c)

	doServerStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Server Streaming RPC")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ana",
			LastName:  "Niki",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC:%v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while readind stream:%v", err)
		}
		log.Printf("Response from GreetManyTimes:%v", msg.GetResult())
	}
}
