package config

import (
	"app/utils"

	"github.com/kelseyhightower/envconfig"
)

type MessagingConfig struct {
	AmpqUrl string `envconfig:"AMPQ_URL"`
}

type ServicesURLs struct {
	PatientServiceUrl    string `envconfig:"PATIENT_SERVICE_URL"`
	DoctorServiceUrl     string `envconfig:"DOCTOR_SERVICE_URL"`
	AppoinmentServiceUrl string `envconfig:"APPOINMENT_SERVICE_URL"`
}

func IniatilizeServicesURLS() *ServicesURLs {
	var s ServicesURLs
	err := envconfig.Process("", &s)
	utils.HandleError(err)
	return &s
}

func IniatilizeMessagingConfig() *MessagingConfig {
	var m MessagingConfig
	err := envconfig.Process("", &m)
	utils.HandleError(err)
	return &m
}
