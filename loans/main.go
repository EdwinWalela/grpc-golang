package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grpc-demo/protos/userproto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:7000"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("failed to connect to grpc server: %v", err)

	}

	defer conn.Close()

	client := pb.NewUserManagmentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	userReq := &pb.UserReq{
		Id: 1,
	}

	user, err := client.GetUser(ctx, userReq)

	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	log.Printf("Recieved user: id:%d name:%s", user.GetId(), user.GetName())

}
