package main

import (
	"log"

	"github.com/atharvamhaske/go-ec2/internal/config"
	"github.com/atharvamhaske/go-ec2/internal/server"
)

func main() {
	cfg := config.Load()

	srv := server.New(cfg)

	log.Printf("server running on port %s", cfg.Port)
	log.Fatal(srv.Start())
}