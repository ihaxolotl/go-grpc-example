package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/ihaxolotl/go-grpc-example/internal/proto"
	"github.com/ihaxolotl/go-grpc-example/internal/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port int = 9999

func main() {
	var (
		listener net.Listener
		server   *grpc.Server
		err      error
	)

	listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	server = grpc.NewServer()
	pb.RegisterUsersServer(server, &transport.UsersServer{})
	reflection.Register(server)
	log.Printf("server listening at %v\n", listener.Addr())

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
