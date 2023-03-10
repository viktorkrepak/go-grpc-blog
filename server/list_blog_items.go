package main

import (
	"context"
	"log"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Printf("ListBlog function was invoked with %v", in)

	cur, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return err
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return err
		}

		stream.Send(docToBlogItem(data))
	}

	if err := cur.Err(); err != nil {
		return err
	}

	return nil
}
