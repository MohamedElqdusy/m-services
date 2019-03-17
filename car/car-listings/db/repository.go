package db

import (
	"context"
	"car-listings/models"
)

type Repository interface {
	StoreCar(ctx context.Context, car models.Car, dealer_id uint64) error
	FindCarsByMake(ctx context.Context, make string) ([]models.Car, error)
	FindCarsByModel(ctx context.Context, model string) ([]models.Car, error)
	FindCarsByYear(ctx context.Context, year uint64) ([]models.Car, error)
	FindCarsByColor(ctx context.Context, color string) ([]models.Car, error)
	FindAllCars(ctx context.Context) ([]models.Car, error)
	Close()
}

var repositoryImpl Repository

func StoreCar(ctx context.Context, car models.Car, dealer_id uint64) error {
	return repositoryImpl.StoreCar(ctx, car, dealer_id)
}

func FindCarsByMake(ctx context.Context, make string) ([]models.Car, error) {
	return repositoryImpl.FindCarsByMake(ctx, make)
}

func FindCarsByModel(ctx context.Context, model string) ([]models.Car, error){
	return repositoryImpl.FindCarsByModel(ctx, model)
}

func FindCarsByYear(ctx context.Context, year uint64) ([]models.Car, error){
	return repositoryImpl.FindCarsByYear(ctx, year)
}

func FindCarsByColor(ctx context.Context, color string) ([]models.Car, error){
	return repositoryImpl.FindCarsByColor(ctx, color)
}

func FindAllCars(ctx context.Context) ([]models.Car, error) {
	return repositoryImpl.FindAllCars(ctx)
}

func Close() {
	repositoryImpl.Close()
}

func SetRepository(repository Repository) {
	repositoryImpl = repository
}
