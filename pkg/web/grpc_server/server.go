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

func (server *Server) Start(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	server.listener = listener

	return server.grpcServer.Serve(listener)
}

func (server *Server) Close() error {
	server.grpcServer.GracefulStop()
	if server.listener == nil {
		return nil
	}

	return server.listener.Close()
}

func (server *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	server.grpcServer.RegisterService(desc, impl)
}
