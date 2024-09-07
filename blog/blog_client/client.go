package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ananikitina/grpc-go/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Blog client")

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.NewClient("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect:%v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Ana",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("unexpected error:%v", err)
	}

	fmt.Printf("Blog has been created:%v", createBlogRes)
	blogId := createBlogRes.GetBlog().GetId()

	//read blog
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "fgdhfghd"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v \n", err2)
	}
	realBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), realBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading:%v \n", readBlogErr)
	}
	fmt.Printf("Blog was read:%v \n", readBlogRes)

	//update blog
	fmt.Println("Updating the blog")

	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Changed Author",
		Title:    "My first blog(edited)",
		Content:  "Content of the first blog with some additions",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("Error happened while updating:%v \n", updateErr)
	}
	fmt.Printf("Blog was updated:%v \n", updateRes)

	//delete blog
	fmt.Println("Deleting the blog")

	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogId})
	if deleteErr != nil {
		fmt.Printf("Error happened while udeleting:%v \n", deleteErr)
	}
	fmt.Printf("Blog was deleted:%v \n", deleteRes)

	//list blogs
	fmt.Println("List blogs")

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC:%v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}

}
