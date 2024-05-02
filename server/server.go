package main

import (
	"context"
	"log"
	"net"

	proto "gosabaq/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedUserServiceServer
}

var users []*proto.User

func (s *server) AddUser(ctx context.Context, user *proto.User) (*proto.User, error) {
	users = append(users, user)
	return user, nil
}

func (s *server) GetUser(ctx context.Context, userID *proto.UserID) (*proto.User, error) {
	for _, u := range users {
		if u.Id == userID.Id {
			return u, nil
		}
	}
	return nil, nil // Returning nil,nil to indicate user not found
}

func (s *server) ListUsers(ctx context.Context, _ *proto.Empty) (*proto.UserList, error) {
	return &proto.UserList{Users: users}, nil
}

func startServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &server{})
	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	startServer()
}
