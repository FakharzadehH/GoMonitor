package job

import (
	"github.com/FakharzadehH/GoMonitor/internal/config"
	"github.com/FakharzadehH/GoMonitor/internal/logger"
	"github.com/FakharzadehH/GoMonitor/repository"
	"github.com/go-co-op/gocron"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"time"
)

func CheckServersHealthJob() {
	cfg := config.GetConfig()
	writeDB, err := config.NewGORMConnection(config.GetConfig().DB.GetWriteDSN())
	if err != nil {
		logger.Logger().Fatal(err)
	}
	readDB, err := config.NewGORMConnection(config.GetConfig().DB.GetReadDSN())
	if err != nil {
		logger.Logger().Fatal(err)
	}
	writeRepo := repository.NewRepository(writeDB) // a good practice would be to have 2 separate repo implementations
	ReadRepo := repository.NewRepository(readDB)
	TehranTime, _ := time.LoadLocation("Asia/Tehran")
	s := gocron.NewScheduler(TehranTime)
	s.Every(cfg.CheckInterval).Minute().Do(func() {
		logger.Logger().Debugw("started Check Servers Health Job")
		//TODO: add metrics of job starts
		startTime := time.Now()
		checkServersHealth(writeRepo, ReadRepo)
		duration := time.Since(startTime)
		logger.Logger().Debugw("Finished Check Servers Health Job", "duration", duration)
	})
	s.StartAsync()
}

func checkServersHealth(writeRepo *repository.Repository, readRepo *repository.Repository) {
	servers, err := readRepo.GetAllServers()
	if err != nil {
		logger.Logger().Errorw("error in CheckServersHealthJob", "error", err)
		return
	}
	for _, server := range servers {
		isHealthy := sendGetRequest(server.Address)
		if isHealthy {
			server.Success += 1
			// TODO: add success metric
		} else {
			server.Failure += 1
			server.LastFailure = null.TimeFrom(time.Now())
			// todo: add failure metric
		}
		if err := writeRepo.Upsert(&server); err != nil {
			logger.Logger().Errorw("error while updating server health status", "error", err)
			//TODO: add metric
		}

	}
}
func sendGetRequest(address string) bool {
	isHealthy := false

	resp, err := http.Get("http://" + address)
	if err != nil {
		return false
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		isHealthy = true
	}

	if resp.StatusCode >= 300 && resp.StatusCode < 400 { // if it's redirected, check https
		resp, _ := http.Get("https://" + address)
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			isHealthy = true
		}
	}
	return isHealthy

}
