package service

import (
	"github.com/FakharzadehH/GoMonitor/repository"
)

type Service struct {
	Repos *repository.Repository
}

func NewService(Repos *repository.Repository) *Service {
	return &Service{
		Repos: Repos,
	}
}
