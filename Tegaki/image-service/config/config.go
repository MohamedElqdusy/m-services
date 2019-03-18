package config

import (
	"image-service/utils"

	"github.com/kelseyhightower/envconfig"
)

type RedisConfig struct {
	RedisAdress   string `envconfig:"REDIS_ADRESS"`
	RedisPassword string `envconfig:"REDIS_PASSWORD"`
	RedisDataBase string `envconfig:"REDIS_DATABASE"`
	RedisPort     string `envconfig:"REDIS_PORT"`
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

func IniatilizeRedisConfig() *RedisConfig {
	var r RedisConfig
	err := envconfig.Process("", &r)
	utils.HandleError(err)
	return &r
}
