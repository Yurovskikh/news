package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	AppPort  int
	NatsHost string
	NatsPort int
}

func ReadConfig() *Config {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	return &Config{
		AppPort:  viper.GetInt("APP_PORT"),
		NatsHost: viper.GetString("NATS_HOST"),
		NatsPort: viper.GetInt("NATS_PORT"),
	}
}
