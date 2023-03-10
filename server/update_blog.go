package main

import (
	"context"
	"fmt"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, payload *pb.BlogItem) (*emptypb.Empty, error) {
	fmt.Printf("UpdateBlog request: %v", payload)

	oid, err := primitive.ObjectIDFromHex(payload.GetId())

	if err != nil {
		return nil, err
	}

	data := &BlogItem{
		AuthorID: payload.GetAuthorId(),
		Title:    payload.GetTitle(),
		Content:  payload.GetContent(),
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("no document matched the filter")
	}

	return &emptypb.Empty{}, nil
}
