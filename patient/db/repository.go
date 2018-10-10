package db

import (
	"context"

	"patient/models"
)

type Repository interface {
	RegisterPatient(ctx context.Context, patient models.Patient) error
	FindAllPatients(ctx context.Context) ([]models.Patient, error)
	FindPatient(ctx context.Context, id uint64) (models.Patient, error)
	Close()
}

var repositoryImpl Repository

func RegisterPatient(ctx context.Context, patient models.Patient) error {
	return repositoryImpl.RegisterPatient(ctx, patient)
}

func FindAllPatients(ctx context.Context) ([]models.Patient, error) {
	return repositoryImpl.FindAllPatients(ctx)
}

func FindPatient(ctx context.Context, id uint64) (models.Patient, error) {
	return repositoryImpl.FindPatient(ctx, id)
}

func Close() {
	repositoryImpl.Close()
}

func SetRepository(repository Repository) {
	repositoryImpl = repository
}
