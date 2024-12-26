package helpers

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func GetRedisClient() *redis.Client {
	return RedisClient
}

func InitRedis() error {
	// redisHost := os.Getenv("REDIS_HOST")
	// redisPort := os.Getenv("REDIS_PORT")
	// redisPassword := os.Getenv("REDIS_PASSWORD")
	// redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	// if err != nil {
	// 	return err
	// }
	// RedisClient = redis.NewClient(&redis.Options{
	// 	Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
	// 	Password: redisPassword,
	// 	DB:       redisDB,
	// })
	redisRailwayUrl := os.Getenv("REDIS_RAILWAY_URL")
	redisOption, err := redis.ParseURL(redisRailwayUrl)
	if err != nil {
		return err
	}
	RedisClient = redis.NewClient(redisOption)
	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func DeleteCachedWithKey(cachedKey string) error {
	var cursor uint64
	for {
		keys, cursor, err := RedisClient.Scan(context.Background(), cursor, cachedKey, 10).Result()
		if err != nil {
			errorS := fmt.Sprintf("Error scanning keys: %s", err)
			return errors.New(errorS)
		}
		for _, key := range keys {
			err := RedisClient.Del(context.Background(), key).Err()
			if err != nil {
				errorS := fmt.Sprintf("Error deleting key: %s - %s", key, err)
				return errors.New(errorS)
			} else {
				fmt.Println("Deleted key: ", key)
			}
		}

		if cursor == 0 {
			break
		}
	}
	return nil
}
