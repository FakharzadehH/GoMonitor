package job

import (
	"github.com/FakharzadehH/GoMonitor/internal/config"
	"github.com/FakharzadehH/GoMonitor/internal/logger"
	"github.com/FakharzadehH/GoMonitor/repository"
	"github.com/go-co-op/gocron"
	"net/http"
	"time"
)

func CheckServersHealthJob() {
	cfg := config.GetConfig()
	db, err := config.NewGORMConnection(cfg)
	if err != nil {
		logger.Logger().Fatal(err)
	}
	repo := repository.NewRepository(db)
	TehranTime, _ := time.LoadLocation("Asia/Tehran")
	s := gocron.NewScheduler(TehranTime)
	s.Every(cfg.CheckInterval).Second().Do(func() {
		logger.Logger().Debugw("started Check Servers Health Job")
		//TODO: add metrics of job starts
		startTime := time.Now()
		checkServersHealth(repo)
		duration := time.Since(startTime)
		logger.Logger().Debugw("Finished Check Servers Health Job", "duration", duration)
	})
	s.StartAsync()
}

func checkServersHealth(repo *repository.Repository) {
	servers, err := repo.GetAllServers()
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
			server.LastFailure = time.Now()
			// todo: add failure metric
		}
		if err := repo.Upsert(&server); err != nil {
			logger.Logger().Errorw("error while updating server health status", "error", err)
			//TODO: add metric
		}

	}
}
func sendGetRequest(address string) bool {
	isHealthy := false

	resp, _ := http.Get("http://" + address)
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
