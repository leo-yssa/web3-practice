package server

import (
	"fmt"
	"net/http"
	"web3-practice/internal/controller"
	"web3-practice/internal/middleware"
	"web3-practice/internal/middleware/validator"
)

type Server struct {
	*http.Server
	serve chan error
}

func NewServer(cfg *Config) (*Server, error) {
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	validator.InitValidator()
	rdb, err := newRDB(cfg)
	if err != nil {
		return nil, err
	}
	cache, err := newCache(cfg)
	if err != nil {
		return nil, err
	}
	ctrl := controller.NewController(rdb, cache)
	return &Server{
		Server: &http.Server{
			Addr:    port,
			Handler: middleware.NewGinHandler(rdb, ctrl),
		},
		serve: make(chan error),
	}, nil
}

func (s *Server) Start(args []string) error {
	go func() {
		if err := s.Listen(); err != nil {
			s.serve <- err
		}
	}()
	return <-s.serve
}

func (s *Server) Listen() error {
	return s.ListenAndServe()
}
