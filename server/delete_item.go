package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) DeleteBlog(ctx context.Context, payload *pb.BlogId) (*pb.BlogId, error) {
	log.Printf("DeleteBlog function was invoked with %v", payload)

	oid, err := primitive.ObjectIDFromHex(payload.GetId())

	if err != nil {
		return nil, err
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("no document matched the filter")
	}

	return payload, nil
}
