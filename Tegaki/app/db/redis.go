package db

import (
	"context"
	"encoding/json"
	"tegaki-service/models"
	"tegaki-service/utils"

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

func (c RedisRepository) FindImageRequestStateById(ctx context.Context, imageId string) (models.ImageRequestState, error) {
	var imgState models.ImageRequestState
	reply, err := c.conn.Get("states:" + imageId).Result()
	utils.HandleError(err)
	if err == nil {
		err = json.Unmarshal([]byte(reply), &imgState)
		utils.HandleError(err)
	}
	return imgState, nil
}

func (c RedisRepository) Close() {
	err := c.conn.Close()
	utils.HandleError(err)
}
