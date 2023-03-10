package main

import (
	"context"
	"io"
	"log"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func getAllBlogItems(c pb.BlogServiceClient) {
	log.Println("---getAllBlogItems was invoked---")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlog RPC: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Println(res)
	}

}
