package main

import pb "github.com/viktorkrepak/go-grpc-blog/proto"

type Server struct {
	pb.BlogServiceServer
}
