package handler

import (
	"context"
	"net/http"

	authenpb "github.com/PhuPhuoc/grpc_micro_test-gateway/proto/authen"
	userpb "github.com/PhuPhuoc/grpc_micro_test-gateway/proto/user"
	"github.com/gin-gonic/gin"
)

// APIHandler handles API routes
type APIHandler struct {
	AuthenClient authenpb.AuthServiceClient
	UserClient   userpb.UserServiceClient

	authenpb.UnimplementedAuthServiceServer
	userpb.UnimplementedUserServiceServer
}

// RegisterRequest defines the request payload for registration
type RegisterRequest struct {
	Email    string `json:"email" example:"user@example.com" binding:"required,email"`
	Password string `json:"password" example:"password123" binding:"required"`
	Name     string `json:"name" example:"John Doe" binding:"required"`
}

// RegisterResponse defines the response payload for registration
type RegisterResponse struct {
	Message string `json:"message" example:"User registered successfully"`
	UserID  string `json:"user_id" example:"12345"`
}

// LoginRequest defines the request payload for login
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com" binding:"required,email"`
	Password string `json:"password" example:"password123" binding:"required"`
}

// LoginResponse defines the response payload for login
type LoginResponse struct {
	Token string      `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."` // Example token
	User  interface{} `json:"user"`                                                    // Can be adjusted based on your User struct
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Create a new user with email, password, and name
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		RegisterRequest	true	"Register request"
//	@Success		200				{object}	RegisterResponse
//	@Router			/register [post]
func (h *APIHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &userpb.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	res, err := h.UserClient.Register(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		Message: res.Message,
		UserID:  res.Id,
	})
}

// Login godoc
//
//	@Summary		Login a user
//	@Description	Authenticate user with email and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		LoginRequest	true	"Login request"
//	@Success		200				{object}	LoginResponse
//	@Router			/login [post]
func (h *APIHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &authenpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.AuthenClient.Login(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: res.Token,
		User:  res.User, // Adjust based on the user information in the response
	})
}
