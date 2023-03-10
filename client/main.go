package main

import (
	"fmt"
	"log"

	pb "github.com/viktorkrepak/go-grpc-blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	fmt.Println(id, "was created")

	res := getById(c, id)
	fmt.Println(res)

	updateBlog(c, id)

	fmt.Println("Blog was updated")

	getAllBlogItems(c)

	fmt.Println("All blog items were retrieved")

	deleteBlog(c, id)

	fmt.Println("Blog was deleted")
}
