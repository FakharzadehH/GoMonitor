package config

import "fmt"

type DB struct {
	WriteHost     string `mapstructure:"write_host"`
	WritePort     string `mapstructure:"write_port"`
	WriteUsername string
	WritePassword string
	ReadHost      string `mapstructure:"read_host"`
	ReadPort      string `mapstructure:"read_port"`
	ReadUsername  string
	ReadPassword  string
	DBName        string `mapstructure:"db_name"`
	SSLMode       string `mapstructure:"ssl_mode"`
}

func (db DB) GetWriteDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db.WriteHost, db.WriteUsername, db.WritePassword, db.DBName, db.WritePort, db.SSLMode)
}

func (db DB) GetReadDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db.ReadHost, db.ReadUsername, db.ReadPassword, db.DBName, db.ReadPort, db.SSLMode)
}
