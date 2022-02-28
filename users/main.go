package main

import (
	"context"

	pb "github.com/grpc-demo/protos/userproto"
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

func main() {

}
