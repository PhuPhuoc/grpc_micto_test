package main

import (
	"log"
	"net"

	"github.com/PhuPhuoc/grpc_micro_test-user/internal/handler"
	pb "github.com/PhuPhuoc/grpc_micro_test-user/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &handler.UserHandler{})

	log.Println("User Service running on port 50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
