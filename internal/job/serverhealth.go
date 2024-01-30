package job

import (
	"github.com/FakharzadehH/GoMonitor/internal/config"
	"github.com/FakharzadehH/GoMonitor/internal/logger"
	"github.com/FakharzadehH/GoMonitor/internal/metrics"
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
	monitorMetrics := metrics.GetMetrics().NewMonitorMetrics("GoMonitor")
	dbMetrics := metrics.GetMetrics().NewDBMetrics()
	writeRepo := repository.NewRepository(writeDB, dbMetrics) // a good practice would be to have 2 separate repo implementations
	ReadRepo := repository.NewRepository(readDB, dbMetrics)
	TehranTime, _ := time.LoadLocation("Asia/Tehran")
	s := gocron.NewScheduler(TehranTime)
	s.Every(cfg.CheckInterval).Minute().Do(func() {
		monitorMetrics.JobStart.Inc()
		startTime := time.Now()
		checkServersHealth(writeRepo, ReadRepo, monitorMetrics)
		duration := time.Since(startTime).Milliseconds()
		monitorMetrics.JobDuration.Add(float64(duration))

		logger.Logger().Debugw("Finished Check Servers Health Job", "duration", duration)
	})
	s.StartAsync()
}

func checkServersHealth(writeRepo *repository.Repository, readRepo *repository.Repository, monitorMetrics *metrics.MonitorMetrics) {
	servers, err := readRepo.GetAllServers()
	if err != nil {
		logger.Logger().Errorw("error while getting servers list in CheckServersHealthJob", "error", err)
		return
	}
	httpClient := http.Client{
		Timeout: 15 * time.Second,
	}
	for _, server := range servers {
		isHealthy, duration := sendGetRequest(server.Address, httpClient, monitorMetrics)
		if isHealthy {
			server.Success += 1
			monitorMetrics.CheckHealthStatusSuccess.Inc()
			monitorMetrics.RequestDurationSuccess.Add(float64(duration))
		} else {
			server.Failure += 1
			server.LastFailure = null.TimeFrom(time.Now())
			monitorMetrics.CheckHealthStatusFailure.Inc()
			monitorMetrics.RequestDurationFailure.Add(float64(duration))

		}
		if err := writeRepo.Upsert(&server); err != nil {
			logger.Logger().Errorw("error while updating server health status", "error", err)
		}

	}
}

func sendGetRequest(address string, httpClient http.Client, monitorMetrics *metrics.MonitorMetrics) (bool, int64) {
	isHealthy := false
	start := time.Now()
	resp, err := httpClient.Get("http://" + address)
	duration := time.Since(start)
	monitorMetrics.SentRequest.Inc()
	if err != nil {
		return false, duration.Milliseconds()
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		isHealthy = true
	}

	if resp.StatusCode >= 300 && resp.StatusCode < 400 { // if it's redirected, check https
		start = time.Now()
		resp, _ := http.Get("https://" + address)
		duration = time.Since(start)
		monitorMetrics.SentRequest.Inc()
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			isHealthy = true
		}
	}
	if !isHealthy {
		logger.Logger().Debugw("server " + address + " is not healthy")
	}
	return isHealthy, duration.Milliseconds()

}
