package config

import (
	"github.com/FakharzadehH/GoMonitor/internal/domain"
	"gorm.io/plugin/prometheus"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGORMConnection(DSN string) (*gorm.DB, error) {

	pg := postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	})

	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  logger.Error,    // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)

	gormDB, err := gorm.Open(pg, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   false,
			TablePrefix:   "",
		},
		PrepareStmt:     false,
		Logger:          log,
		CreateBatchSize: 100,
	})
	gormDB.Use(prometheus.New(prometheus.Config{
		DBName:          GetConfig().DB.DBName,
		RefreshInterval: 15,   // refresh metrics interval (default 15 seconds)
		StartServer:     true, // start http server to expose metrics
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.Postgres{VariableNames: []string{"Threads_running"}},
		},
		Labels: map[string]string{
			"instance": "127.0.0.1",
		},
	}))

	if err != nil {
		return nil, err
	}
	if !gormDB.Migrator().HasTable("servers") {
		gormDB.AutoMigrate(&domain.ServerStatus{})
	}
	return gormDB, nil
}
