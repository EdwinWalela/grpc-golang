package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	pb "github.com/grpc-demo/protos/userproto"
	"google.golang.org/grpc"
)

const (
	userGrpcAddress = "localhost:7000"
	muxAddress      = "0.0.0.0:8000"
)

var client pb.UserManagmentClient

func handleNewLoan(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	userReq := &pb.UserReq{
		Id: 1,
	}

	user, err := client.GetUser(ctx, userReq)

	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	res := fmt.Sprintf("Recieved user: id:%d name:%s", user.GetId(), user.GetName())

	json.NewEncoder(w).Encode(res)
}

func main() {

	// GRPC Connection
	conn, err := grpc.Dial(userGrpcAddress, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("failed to connect to grpc server: %v", err)
	}

	defer conn.Close()

	client = pb.NewUserManagmentClient(conn)

	r := mux.NewRouter()
	r.HandleFunc("/", handleNewLoan).Methods("POST")

	srv := http.Server{
		Addr:         muxAddress,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Printf("Listening for requests on: %s", muxAddress)

	log.Fatal(srv.ListenAndServe())

}
