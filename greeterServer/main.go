//go:generate protoc -I ../chat --go_out=plugins=grpc:../chat ../chat/chat.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	pb "goGRPC/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement chat.GreeterServer.
type server struct{}

// SendMessage implements chat.GreeterServer
func (s *server) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Message from %v : %v", in.Name, in.Message)
	return &pb.MessageResponse{Receive: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
