package routes

import (
	"net/http"

	"github.com/atharvamhaske/go-ec2/internal/handlers"
	"github.com/atharvamhaske/go-ec2/internal/services"
)


func Register() http.Handler {
	mux := http.NewServeMux()

	healthService := services.New()
	healthHandler := handlers.New(healthService)

	mux.HandleFunc("/health", healthHandler.Check)

	return mux
}