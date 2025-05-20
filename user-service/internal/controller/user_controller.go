package controller

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user-service/internal/models"
	"user-service/internal/usecase"
	pb "user-service/proto"
)

type UserController struct {
	pb.UnimplementedUserServiceServer
	useCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) *UserController {
	return &UserController{useCase: useCase}
}

func (c *UserController) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.UserResponse, error) {
	if req.GetUsername() == "" || req.GetPassword() == "" || req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "username, password and email are required")
	}

	user := &models.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	}

	if err := c.useCase.RegisterUser(ctx, user); err != nil {
		if errors.Is(err, usecase.ErrUsernameExists) {
			return nil, status.Error(codes.AlreadyExists, "username already exists")
		}
		log.Printf("Registration failed: %v", err)
		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &pb.UserResponse{
		Id:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (c *UserController) AuthenticateUser(ctx context.Context, req *pb.AuthenticateUserRequest) (*pb.AuthResponse, error) {
	if req.GetUsername() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "username and password are required")
	}

	user, err := c.useCase.AuthenticateUser(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			return nil, status.Error(codes.Unauthenticated, "invalid credentials")
		}
		log.Printf("Authentication failed: %v", err)
		return nil, status.Error(codes.Internal, "authentication failed")
	}

	// In production, use a proper JWT token generator
	token := "generated-jwt-token-" + user.ID.Hex()

	return &pb.AuthResponse{
		Token:   token,
		UserId: user.ID.Hex(),
	}, nil
}

func (c *UserController) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.UserResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is required")
	}

	user, err := c.useCase.GetUserProfile(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Printf("Failed to get user profile: %v", err)
		return nil, status.Error(codes.Internal, "failed to get user profile")
	}

	return &pb.UserResponse{
		Id:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (c *UserController) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Status: "SERVING"}, nil
}