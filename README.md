# gRPC demo (Golang)

Microservices communication using [gRPC](https://grpc.io/)

## Generate server and client code

`cd protos/userproto`


`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative users.proto`

## Run services

Start Users gRPC server

`go run /users/main.go` 


Start Loans Mux server

`go run /loans/main.go`


### `/POST localhost:8000`

