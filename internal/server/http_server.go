package server

import (
	"fmt"
	"net/http"
	"web3-practice/internal/config"
	"web3-practice/internal/controller"
	"web3-practice/internal/middleware"
	"web3-practice/internal/middleware/validator"
	"web3-practice/internal/repository"
)

type HttpServer struct {
	*http.Server
	serve chan error
}

func NewHttpServer(cfg *config.Config) (*HttpServer, error) {
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
	repo := repository.NewRepository(rdb)
	if err := repo.Initialize(); err != nil {
		return nil, err
	}
	ctrl := controller.NewController(repo, cache, cfg)
	return &HttpServer{
		Server: &http.Server{
			Addr:    port,
			Handler: middleware.NewGinHandler(repo, ctrl, cfg),
		},
		serve: make(chan error),
	}, nil
}

func (s *HttpServer) Start(args []string) error {
	go func() {
		if err := s.ListenAndServe(); err != nil {
			s.serve <- err
		}
	}()
	return <-s.serve
}
