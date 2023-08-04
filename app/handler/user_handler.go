package handler

import (
	"context"
	"github/lzzzzl/chat-chat-go/domain/service"
	pb "github/lzzzzl/chat-chat-go/infrastructure/protocol"
)

// UserHandler ...
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler ...
func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

// Register ...
func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := h.userService.Register(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// Convert user to protobuf user
	return &pb.RegisterResponse{
		Username: user.Username,
		Success:  true,
		Message:  "",
	}, nil
}
