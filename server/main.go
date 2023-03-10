package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var collection *mongo.Collection

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017"))

	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server listening on %v", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
