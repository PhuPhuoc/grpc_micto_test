package handler

import (
	"context"
	"fmt"

	pb "github.com/PhuPhuoc/grpc_micro_test-user/proto"
)

type UserHandler struct {
	// Add database connection or other dependencies if needed
	pb.UnimplementedUserServiceServer
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.Email == "test@example.com" {
		return nil, fmt.Errorf("email already exists")
	}

	return &pb.RegisterResponse{
		Id:      "123",
		Message: "User registered successfully",
	}, nil
}
