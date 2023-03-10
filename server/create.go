package main

import (
	"context"
	"log"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) CreateBlog(ctx context.Context, payload *pb.BlogItem) (*pb.BlogItem, error) {
	log.Printf("CreateBlog request: %v", payload)

	data := BlogItem{
		AuthorID: payload.AuthorId,
		Title:    payload.GetTitle(),
		Content:  payload.GetContent(),
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, err
	}

	return &pb.BlogItem{
		Id:       oid.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}, nil
}
