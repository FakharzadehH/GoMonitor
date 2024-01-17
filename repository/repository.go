package repository

import (
	"github.com/FakharzadehH/GoMonitor/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Upsert(status *domain.ServerStatus) error {
	return r.db.Save(status).Error
}

func (r *Repository) GetServerStatusByID(id uint, status *domain.ServerStatus) error {
	return r.db.First(status, id).Error
}

func (r *Repository) GetServerStatusByAddress(address string, status *domain.ServerStatus) error {
	return r.db.Where("address = ?", address).First(&status).Error
}

func (r *Repository) GetAllServers() ([]domain.ServerStatus, error) {
	servers := []domain.ServerStatus{}
	return servers, r.db.Find(&servers).Error
}
