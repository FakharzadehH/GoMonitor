package config

import (
	"github.com/spf13/viper"
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

	return nil
}
