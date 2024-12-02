package main

import (
	"fmt"
	"log"

	docs "github.com/PhuPhuoc/grpc_micro_test-gateway/docs"
	"github.com/PhuPhuoc/grpc_micro_test-gateway/internal/handler"
	authenpb "github.com/PhuPhuoc/grpc_micro_test-gateway/proto/authen"
	userpb "github.com/PhuPhuoc/grpc_micro_test-gateway/proto/user"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

//	@title		Demo gRPC microservices - Gateway service
//	@version	1.0

func main() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	// Kết nối tới Authen Service
	authenConn, err := grpc.Dial("authen-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Authen Service: %v", err)
	}
	defer authenConn.Close()

	// Kết nối tới User Service
	userConn, err := grpc.Dial("user-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to User Service: %v", err)
	}
	defer userConn.Close()

	// Tạo các client gRPC
	authenClient := authenpb.NewAuthServiceClient(authenConn)
	userClient := userpb.NewUserServiceClient(userConn)

	// Khởi tạo router Gin
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Sử dụng middleware CORS
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/swagger/*any"),
		gin.Recovery(),
	)
	apiHandler := &handler.APIHandler{
		AuthenClient: authenClient,
		UserClient:   userClient,
	}

	// Định nghĩa các endpoint
	router.POST("/api/v1/register", apiHandler.Register)
	router.POST("/api/v1/login", apiHandler.Login)

	// Chạy Gateway
	log.Println("API Gateway running on port 8080...")
	fmt.Printf("\nFor development: http://localhost%s/swagger/index.html\n\n", "8080")
	router.Run(":8080")
}
