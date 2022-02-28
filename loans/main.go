package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	pb "github.com/grpc-demo/protos/userproto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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

	err := godotenv.Load()

	userGrpcAddress := os.Getenv("users_grpc_addr")
	muxAddress := os.Getenv("users_service_addr")

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	// GRPC Connection
	conn, err := grpc.Dial(userGrpcAddress, grpc.WithInsecure())

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
