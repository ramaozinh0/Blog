package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ramaozinh0/grpcCourse/blog/blogpb"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() // Maybe this should be in a separate function and the error handled?

	c := blogpb.NewBlogServiceClient(cc)

	// create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Ramon alves",
		Title:    "Meu primeiro blog",
		Content:  "Conteudo do blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createBlogRes)
	blogID:= createBlogRes.GetBlog().GetId()
	//blogID := createBlogRes.GetBlog().GetId()
	


	// READ BLOG
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "asdadsad"})
	if err2 != nil{
		fmt.Printf("Error happened while reading: %v \n", err2)
	}
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil{
		fmt.Printf("Error happened while reading: %v \n", readBlogErr)
	}
	fmt.Printf("Blog was read: %v\n", readBlogRes)

	//update blog
	newBlog:= &blogpb.Blog{
		Id: blogID,
		AuthorId: "Changed author",
		Title: "Maicoms sales",
		Content: "Um blog legalzao",
	}
	updateRes, updateErr:=c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr!= nil{
		log.Printf("Error while updating: %v\n", updateErr)
	}
	fmt.Printf("Blog updated: %v", updateRes)
}