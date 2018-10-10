package db

import (
	"context"

	"appoinment/models"
)

type Repository interface {
	ReserveAppoinment(ctx context.Context, appoinment models.Appoinment) error
	FindAllDoctorAppoinments(ctx context.Context, doctorId int) ([]models.Appoinment, error)
	Close()
}

var repositoryImpl Repository

func ReserveAppoinment(ctx context.Context, appoinment models.Appoinment) error {
	return repositoryImpl.ReserveAppoinment(ctx, appoinment)
}

func FindAllDoctorAppoinments(ctx context.Context, doctorId int) ([]models.Appoinment, error) {
	return repositoryImpl.FindAllDoctorAppoinments(ctx, doctorId)
}

func Close() {
	repositoryImpl.Close()
}

func SetRepository(repository Repository) {
	repositoryImpl = repository
}
