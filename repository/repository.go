package repository

import (
	"github.com/FakharzadehH/GoMonitor/internal/domain"
	"github.com/FakharzadehH/GoMonitor/internal/metrics"
	"gorm.io/gorm"
)

type Repository struct {
	db      *gorm.DB
	metrics *metrics.DBMetrics
}

func NewRepository(db *gorm.DB, metrics *metrics.DBMetrics) *Repository {
	return &Repository{
		db:      db,
		metrics: metrics,
	}
}

func (r *Repository) Upsert(status *domain.ServerStatus) error {
	err := r.db.Save(status).Error
	if err != nil {
		r.metrics.DBReplyFailure.Inc()
	}
	return err
}

func (r *Repository) GetServerStatusByID(id uint, status *domain.ServerStatus) error {
	err := r.db.First(status, id).Error
	if err != nil {
		r.metrics.DBReplyFailure.Inc()
	}
	return err
}

func (r *Repository) GetServerStatusByAddress(address string, status *domain.ServerStatus) error {
	err := r.db.Where("address = ?", address).First(&status).Error
	if err != nil {
		r.metrics.DBReplyFailure.Inc()
	}
	return err
}

func (r *Repository) GetAllServers() ([]domain.ServerStatus, error) {
	servers := []domain.ServerStatus{}
	err := r.db.Find(&servers).Error
	if err != nil {
		r.metrics.DBReplyFailure.Inc()
	}
	return servers, err
}
