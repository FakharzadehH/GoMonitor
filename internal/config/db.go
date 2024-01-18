package config

import "fmt"

type DB struct {
	WriteHost     string `mapstructure:"write_host"`
	WritePort     string `mapstructure:"write_port"`
	WriteUsername string `mapstructure:"write_username"`
	WritePassword string `mapstructure:"write_password"`
	ReadHost      string `mapstructure:"read_host"`
	ReadPort      string `mapstructure:"read_port"`
	ReadUsername  string `mapstructure:"read_username"`
	ReadPassword  string `mapstructure:"read_password"`
	DBName        string `mapstructure:"db_name"`
	SSLMode       string `mapstructure:"ssl_mode"`
}

const DBMSPostgres = "postgres"

func (db DB) GetWriteDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db.WriteHost, db.WriteUsername, db.WritePassword, db.DBName, db.WritePort, db.SSLMode)
}

func (db DB) GetReadDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db.ReadHost, db.ReadUsername, db.ReadPassword, db.DBName, db.ReadPort, db.SSLMode)
}
