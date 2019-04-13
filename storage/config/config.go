package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	NatsHost         string
	NatsPort         int
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

func ReadConfig() *Config {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	return &Config{
		NatsHost:         viper.GetString("NATS_HOST"),
		NatsPort:         viper.GetInt("NATS_PORT"),
		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresPort:     viper.GetInt("POSTGRES_PORT"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),
		PostgresDBName:   viper.GetString("POSTGRES_DBNAME"),
	}
}
