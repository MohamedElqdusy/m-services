package db

import (
	"context"
	"encoding/json"
	"image-service/models"
	"image-service/utils"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	conn *redis.Client
}

// Connecting to Redis server
func RedisConnect(url string, password string, database int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       database,
	})
	return client
}

func NewRedis(url string, password string, database int) (*RedisRepository, error) {
	return &RedisRepository{
		RedisConnect(url, password, database),
	}, nil
}

func (c RedisRepository) StoreImageRequestState(ctx context.Context, s models.ImageRequestState) error {
	blob, err := json.Marshal(s)
	// Save JSON to redis
	err = c.conn.Set("states:"+s.Id, blob, 0).Err()
	utils.HandleError(err)
	return nil
}

func (c RedisRepository) Close() {
	err := c.conn.Close()
	utils.HandleError(err)
}
