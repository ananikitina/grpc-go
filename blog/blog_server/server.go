package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ananikitina/grpc-go/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

func main() {
	//if we crash we get the file name and lne number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve:%v", err)
		}
	}()

	//Wait until a signal is recieved
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is recievd
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("End of program")

}
