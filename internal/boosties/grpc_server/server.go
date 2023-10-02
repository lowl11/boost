package grpc_server

import (
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func New(options ...grpc.ServerOption) *Server {
	return &Server{
		grpcServer: grpc.NewServer(options...),
	}
}
