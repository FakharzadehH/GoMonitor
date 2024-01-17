package domain

import "time"

type ServerStatus struct {
	ID          uint      `json:"id"`
	Address     string    `json:"address"`
	Success     int       `json:"success"`
	Failure     int       `json:"failure"`
	LastFailure time.Time `json:"last_failure"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ServerStatus) TableName() string {
	return "servers"
}

type AddServerRequest struct {
	Address string `json:"address"`
}
type AddServerResponse struct {
	ServerID uint `json:"server_id"`
}

type StatusShowResponse struct {
	Server ServerStatus `json:"server"`
}

type StatusIndexResponse struct {
	Servers []ServerStatus `json:"servers"`
}
