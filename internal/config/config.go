package config

import (
	"github.com/spf13/viper"
	"strings"
)

var cfg Config

type Config struct {
	ReadDB        DB     `mapstructure:"read_db"`
	WriteDB       DB     `mapstructure:"write_db"`
	CheckInterval string `mapstructure:"check_interval"`
	AppPort       string `mapstructure:"app_port"`
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

	return nil
}
