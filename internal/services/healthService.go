package services

import "github.com/atharvamhaske/go-ec2/internal/models"

type HealthService struct {}

func New() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Status() models.Health {
	return models.Health{
		Status: "UP and running",
	}
}