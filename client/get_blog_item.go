package main

import (
	"context"
	"log"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
)

func getById(c pb.BlogServiceClient, id string) *pb.BlogItem {
	log.Println("---getById was invoked---")

	blogId := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), blogId)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
