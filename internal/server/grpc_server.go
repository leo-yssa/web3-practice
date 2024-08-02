package server

import (
	"fmt"
	"net"
	"web3-practice/internal/config"

	"google.golang.org/grpc"
)

const GrpcServerName = "gateway"

type GrpcServer struct {
	*grpc.Server
	listener net.Listener
	serve    chan error
}

func NewGrpcServer(cfg *config.Config) (*GrpcServer, error) {
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()

	return &GrpcServer{
		Server:   s,
		listener: listener,
		serve:    make(chan error),
	}, nil
}

func (s *GrpcServer) Start(args []string) error {
	go func() {
		if err := s.Serve(s.listener); err != nil {
			s.serve <- err
		}
	}()
	return <-s.serve
}
