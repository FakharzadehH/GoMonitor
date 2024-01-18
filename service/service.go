package service

import (
	"github.com/FakharzadehH/GoMonitor/internal/domain"
	"github.com/FakharzadehH/GoMonitor/internal/logger"
	"github.com/FakharzadehH/GoMonitor/repository"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Service struct {
	writeRepo *repository.Repository
	readRepo  *repository.Repository
}

func NewService(writeRepo *repository.Repository, readRepo *repository.Repository) *Service {
	return &Service{
		writeRepo: writeRepo,
		readRepo:  readRepo,
	}
}

func (s *Service) AddServer(payload domain.AddServerRequest) (*domain.AddServerResponse, error) {
	serverStatus := &domain.ServerStatus{}
	err := s.readRepo.GetServerStatusByAddress(payload.Address, serverStatus)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			logger.Logger().Errorw("error while adding server", "error", err)
			return nil, err
		}
		return nil, errors.New("A server with this address is already registered")
	}
	serverStatus.Address = payload.Address
	if err := s.writeRepo.Upsert(serverStatus); err != nil {
		return nil, err
	}
	return &domain.AddServerResponse{ServerID: serverStatus.ID}, nil
}

func (s *Service) GetServerByID(id uint) (*domain.StatusShowResponse, error) {
	serverStatus := domain.ServerStatus{}
	if err := s.readRepo.GetServerStatusByID(id, &serverStatus); err != nil {
		logger.Logger().Errorw("Error while getting server from db by id ", "error", err)
		return nil, err
	}
	return &domain.StatusShowResponse{Server: serverStatus}, nil
}

func (s *Service) GetAllServers() (*domain.StatusIndexResponse, error) {
	servers, err := s.readRepo.GetAllServers()
	if err != nil {
		logger.Logger().Errorw("error while getting all servers", "error", err)
		return nil, err
	}
	return &domain.StatusIndexResponse{Servers: servers}, nil
}
