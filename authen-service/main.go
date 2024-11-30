package main

import (
	"log"
	"net"

	"github.com/PhuPhuoc/grpc_micro_test-authen/internal/handler"
	pd "github.com/PhuPhuoc/grpc_micro_test-authen/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pd.RegisterAuthServiceServer(s, &handler.AuthenHandler{})

	log.Println("authen server runnint on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
