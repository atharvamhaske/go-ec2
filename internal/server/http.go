package server

import (
	"net/http"

	"github.com/atharvamhaske/go-ec2/internal/config"
	"github.com/atharvamhaske/go-ec2/internal/routes"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	handler := routes.Register()
	return http.ListenAndServe(s.cfg.Port, handler)
}