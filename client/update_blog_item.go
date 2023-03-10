package main

import (
	"context"
	"log"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	updatedItem := &pb.BlogItem{
		Id:       id,
		AuthorId: "Viktor",
		Title:    "Updated blog",
		Content:  "Updated content of the test blog",
	}

	_, err := c.UpdateBlog(context.Background(), updatedItem)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was updated: %v\n", updatedItem)
}
