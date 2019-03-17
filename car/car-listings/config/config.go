package config

import (
	"car-listings/utils"
	"github.com/kelseyhightower/envconfig"
)

type PostgreConfig struct {
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	PostgresDataBase string `envconfig:"POSTGRES_DATABASE"`
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     string `envconfig:"POSTGRES_PORT"`
}

func IniatilizePostgreConfig() *PostgreConfig {
	var p PostgreConfig
	err := envconfig.Process("", &p)
	utils.HandleError(err)
	return &p
}
