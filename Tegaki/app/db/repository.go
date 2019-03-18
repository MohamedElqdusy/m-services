package db

import (
	"context"

	"tegaki-service/models"
)

type Repository interface {
	FindImageRequestStateById(ctx context.Context, imageId string) (models.ImageRequestState, error)
	Close()
}

var repositoryImpl Repository

func SetRepository(repository Repository) {
	repositoryImpl = repository
}

func FindImageRequestStateById(ctx context.Context, imageId string) (models.ImageRequestState, error) {
	return repositoryImpl.FindImageRequestStateById(ctx, imageId)

}

func Close() {
	repositoryImpl.Close()
}
