package config

import (
	"doctor/utils"

	"github.com/kelseyhightower/envconfig"
)

type PostgreConfig struct {
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	PostgresDataBase string `envconfig:"POSTGRES_DATABASE"`
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     string `envconfig:"POSTGRES_PORT"`
}

type MessagingConfig struct {
	AmpqUrl string `envconfig:"AMPQ_URL"`
}

func IniatilizeMessagingConfig() *MessagingConfig {
	var m MessagingConfig
	err := envconfig.Process("", &m)
	utils.HandleError(err)
	return &m
}

func IniatilizePostgreConfig() *PostgreConfig {
	var p PostgreConfig
	err := envconfig.Process("", &p)
	utils.HandleError(err)
	return &p
}
