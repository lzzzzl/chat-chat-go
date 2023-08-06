package main

import (
	"context"
	"log"
	"net"

	pb "github/lzzzzl/chat-chat-go/infrastructure/protocol"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) SendMessage(ctx context.Context, req *pb.ChatRequest) (*pb.ChatResponse, error) {
	log.Printf("Received message from %s: %s", req.User, req.Message)
	return &pb.ChatResponse{User: "Server", Message: "Message received"}, nil
}

func main() {
	// 创建gRPC服务
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterChatServiceServer(s, &server{})

	// 侦听端口
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 启动服务器
	log.Println("Server is running on localhost:50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
