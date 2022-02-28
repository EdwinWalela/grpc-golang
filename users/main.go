package main

import (
	"context"
	"log"
	"net"

	pb "github.com/grpc-demo/protos/userproto"
	"google.golang.org/grpc"
)

type User struct {
	id    int32
	name  string
	email string
}

var users = []User{
	{
		id:    "1",
		name:  "edwin",
		email: "edwin@mail.com",
	},
	{
		id:    "2",
		name:  "edwin",
		email: "edwin@mail.com",
	},
	{
		id:    "3",
		name:  "edwin",
		email: "edwin@mail.com",
	},
	{
		id:    "4",
		name:  "edwin",
		email: "edwin@mail.com",
	},
}

type UserManagementServer struct {
	pb.UnimplementedUserManagmentServer
}

func (s *UserManagementServer) GetUser(ctx context.Context, userReq *pb.UserReq) (*pb.User, error) {
	res := users[userReq.GetId()]

	user := &pb.User{
		Id:    res.id,
		Email: res.email,
		Name:  res.name,
	}

	return user, nil
}

const (
	port = ":7000"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserManagmentServer(server, &UserManagementServer{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
