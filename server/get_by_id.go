package main

import (
	"context"
	"log"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) ReadBlog(ctx context.Context, payload *pb.BlogId) (*pb.BlogItem, error) {
	log.Printf("ReadBlog request: %v", payload)

	oid, err := primitive.ObjectIDFromHex(payload.GetId())

	if err != nil {
		return nil, err
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	err = collection.FindOne(ctx, filter).Decode(data)

	if err != nil {
		return nil, err
	}

	return docToBlogItem(data), nil
}
