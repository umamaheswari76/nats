package main

import (
	"fmt"
	pro "kv_grpc/proto"
	"kv_grpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pro.RegisterKvGrpcServiceServer(s, &service.Server{})

	fmt.Println("Server is running on : 50051")
	
	if err := s.Serve(listen); err!=nil{
		log.Fatalf("Failed to serve: %v", err)
	}
}
