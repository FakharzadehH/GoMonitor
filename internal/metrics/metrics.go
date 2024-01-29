package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	sentRequest       *prometheus.CounterVec
	checkHealthStatus *prometheus.CounterVec
	requestDuration   *prometheus.CounterVec
	dbReply           *prometheus.CounterVec
	jobDuration       *prometheus.CounterVec
	jobStart          *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		sentRequest: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "sent_request",
			Help: "number of sent requests",
		}, []string{"job_name"}),
		checkHealthStatus: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "check_health_status",
			Help: "shows servers health check result",
		}, []string{"job_name", "status"}),
		requestDuration: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "check_health_request_duration",
			Help: "duration it takes for getting a response",
		}, []string{"job_name", "status"}),
		dbReply: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "db_reply",
			Help: "status of db replies to queries",
		}, []string{"db", "status"}),
		jobDuration: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "job_duration",
			Help: "duration it takes to finish checking all servers healths",
		}, []string{"job_name"}),
		jobStart: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "job_start",
			Help: "number of job starts",
		}, []string{"job_name"}),
	}
}

type MonitorMetrics struct {
	SentRequest              prometheus.Counter
	CheckHealthStatusSuccess prometheus.Counter
	CheckHealthStatusFailure prometheus.Counter
	RequestDurationSuccess   prometheus.Counter
	RequestDurationFailure   prometheus.Counter
	DBReplyFailure           prometheus.Counter
	JobDuration              prometheus.Counter
	JobStart                 prometheus.Counter
}

func (m *Metrics) NewMonitorMetrics(name string) *MonitorMetrics {
	return &MonitorMetrics{
		SentRequest:              m.sentRequest.WithLabelValues(name),
		CheckHealthStatusSuccess: m.checkHealthStatus.WithLabelValues(name, "success"),
		CheckHealthStatusFailure: m.checkHealthStatus.WithLabelValues(name, "failure"),
		RequestDurationSuccess:   m.requestDuration.WithLabelValues(name, "success"),
		RequestDurationFailure:   m.requestDuration.WithLabelValues(name, "failure"),
		DBReplyFailure:           m.dbReply.WithLabelValues("gomonitor", "failure"),
		JobDuration:              m.jobDuration.WithLabelValues(name),
		JobStart:                 m.jobStart.WithLabelValues(name),
	}
}
