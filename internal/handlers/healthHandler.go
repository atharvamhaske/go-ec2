package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/atharvamhaske/go-ec2/internal/services"
)

type HealthHandler struct {
	service *services.HealthService
}

func New(service *services.HealthService) *HealthHandler {
	return &HealthHandler{
		service: service,
	}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	response := h.service.Status()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}