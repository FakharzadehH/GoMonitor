package repository

import (
	"github.com/FakharzadehH/GoMonitor/internal/domain"
	"github.com/FakharzadehH/GoMonitor/internal/metrics"
	"gorm.io/gorm"
	"time"
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
	start := time.Now()
	err := r.db.Save(status).Error
	duration := time.Since(start).Milliseconds()
	r.metrics.DBCalls.Inc()
	r.metrics.DBLatency.Add(float64(duration))
	if err != nil {
		r.metrics.DBReplyFailure.Inc()
	}
	return err
}

func (r *Repository) GetServerStatusByID(id uint, status *domain.ServerStatus) error {
	start := time.Now()
	err := r.db.First(status, id).Error
	duration := time.Since(start).Milliseconds()

	r.metrics.DBCalls.Inc()
	r.metrics.DBLatency.Add(float64(duration))

	if err != nil {
		r.metrics.DBReplyFailure.Inc()
	}
	return err
}

func (r *Repository) GetServerStatusByAddress(address string, status *domain.ServerStatus) error {
	start := time.Now()
	err := r.db.Where("address = ?", address).First(&status).Error
	duration := time.Since(start).Milliseconds()

	r.metrics.DBCalls.Inc()
	r.metrics.DBLatency.Add(float64(duration))

	if err != nil {
		r.metrics.DBReplyFailure.Inc()
	}
	return err
}

func (r *Repository) GetAllServers() ([]domain.ServerStatus, error) {
	servers := []domain.ServerStatus{}
	start := time.Now()
	err := r.db.Find(&servers).Error
	duration := time.Since(start).Milliseconds()

	r.metrics.DBCalls.Inc()
	r.metrics.DBLatency.Add(float64(duration))

	if err != nil {
		r.metrics.DBReplyFailure.Inc()
	}
	return servers, err
}
