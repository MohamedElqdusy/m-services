package db

import (
	"context"

	"doctor/models"
)

type Repository interface {
	RegisterDoctor(ctx context.Context, doctor models.Doctor) error
	FindAllDoctors(ctx context.Context) ([]models.Doctor, error)
	FindDoctor(ctx context.Context, id uint64) (models.Doctor, error)
	Close()
}

var repositoryImpl Repository

func RegisterDoctor(ctx context.Context, doctor models.Doctor) error {
	return repositoryImpl.RegisterDoctor(ctx, doctor)
}

func FindAllDoctors(ctx context.Context) ([]models.Doctor, error) {
	return repositoryImpl.FindAllDoctors(ctx)
}

func FindDoctor(ctx context.Context, id uint64) (models.Doctor, error) {
	return repositoryImpl.FindDoctor(ctx, id)
}

func Close() {
	repositoryImpl.Close()
}

func SetRepository(repository Repository) {
	repositoryImpl = repository
}
