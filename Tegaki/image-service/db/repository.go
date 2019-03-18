package db

import (
	"context"

	"image-service/models"
)

type Repository interface {
	StoreImageRequestState(ctx context.Context, imageRequestState models.ImageRequestState) error
	Close()
}

var repositoryImpl Repository

func SetRepository(repository Repository) {
	repositoryImpl = repository
}

func StoreImageRequestState(ctx context.Context, imageRequestState models.ImageRequestState) error {
	return repositoryImpl.StoreImageRequestState(ctx, imageRequestState)
}

func Close() {
	repositoryImpl.Close()
}
