package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfg Config

type Config struct {
	DB            DB  `mapstructure:"db"`
	CheckInterval int `mapstructure:"check_interval"`
}

func GetConfig() Config {
	return cfg
}

func Load(configPath string) error {
	v := viper.New()
	v.SetEnvPrefix("CC_Project")
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	//writeUsername, _ := b64.StdEncoding.DecodeString(os.Getenv("DB_WRITE_USERNAME"))
	//writePassword, _ := b64.StdEncoding.DecodeString(os.Getenv("DB_WRITE_PASSWORD"))
	//readUsername, _ := b64.StdEncoding.DecodeString(os.Getenv("DB_READ_USERNAME"))
	//readPassword, _ := b64.StdEncoding.DecodeString(os.Getenv("DB_WRITE_PASSWORD"))
	//cfg.DB.WriteUsername = string(writeUsername)
	//cfg.DB.WritePassword = string(writePassword)
	//cfg.DB.ReadUsername = string(readUsername)
	//cfg.DB.ReadPassword = string(readPassword)
	writeUsername := os.Getenv("DB_WRITE_USERNAME")
	writePassword := os.Getenv("DB_WRITE_PASSWORD")
	readUsername := os.Getenv("DB_READ_USERNAME")
	readPassword := os.Getenv("DB_WRITE_PASSWORD")
	cfg.DB.WriteUsername = writeUsername
	cfg.DB.WritePassword = writePassword
	cfg.DB.ReadUsername = readUsername
	cfg.DB.ReadPassword = readPassword
	return nil
}
