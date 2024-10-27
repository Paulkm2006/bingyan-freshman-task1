package utils

import (
	"bingyan-freshman-task0/internal/config"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func WriteValidationCode(email string, code string) error {
	ctx := context.Background()
	err := RedisClient.Set(ctx, email, code, 0).Err()
	if err != nil {
		return err
	}
	err = RedisClient.Expire(ctx, email, time.Duration(config.Config.Mail.Expire*60)*time.Second).Err()
	return err
}
func GetValidationCode(email string) (string, error) {
	ctx := context.Background()
	code, err := RedisClient.Get(ctx, email).Result()
	return code, err
}

func ValidateCode(email string, code string) (bool, error) {
	ctx := context.Background()
	codeInRedis, err := RedisClient.Get(ctx, email).Result()
	if err != nil {
		return false, err
	}
	if codeInRedis == code {
		return true, nil
	}
	return false, nil
}

func CheckEmailExist(email string) (bool, error) {
	ctx := context.Background()
	_, err := RedisClient.Get(ctx, email).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
