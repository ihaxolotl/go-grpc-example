package transport

import (
	"context"

	pb "github.com/ihaxolotl/go-grpc-example/internal/proto"
)

type UsersServer struct {
	pb.UnimplementedUsersServer
}

func (srv *UsersServer) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}

func (srv *UsersServer) GetUsers(
	ctx context.Context,
	req *pb.GetUsersRequest,
) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{}, nil
}

func (srv *UsersServer) GetUserById(
	ctx context.Context,
	req *pb.GetUserByIdRequest,
) (*pb.GetUserByIdResponse, error) {
	return &pb.GetUserByIdResponse{}, nil
}

func (srv *UsersServer) DeleteUser(
	ctx context.Context,
	req *pb.DeleteUserRequest,
) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{}, nil
}
