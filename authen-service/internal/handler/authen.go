package handler

import (
	"context"
	"fmt"
	"log"

	pd "github.com/PhuPhuoc/grpc_micro_test-authen/proto"

)

/*
UnimplementedAuthServiceServer là một struct được gRPC sinh ra,
cung cấp các phương thức mặc định (stub) cho tất cả các phương thức trong service.

Bằng cách nhúng (embed) struct này vào AuthenHandler,
bạn không cần phải tự triển khai tất cả các phương thức, chỉ cần tập trung vào những gì bạn cần.
*/
type AuthenHandler struct {
	// add db connection or other dependencies if needed
	pd.UnimplementedAuthServiceServer
}

func (h *AuthenHandler) Login(ctx context.Context, req *pd.LoginRequest) (*pd.LoginResponse, error) {
	log.Println("req: ", req)

	if req.Email == "test@gmail.com" && req.Password == "123" {
		return &pd.LoginResponse{
			Token: "dummy-jwt-token",
			User: &pd.UserInfo{
				Id:    "1",
				Email: "test@gmail.com",
				Name:  "Phu Phuoc",
			},
		}, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}
