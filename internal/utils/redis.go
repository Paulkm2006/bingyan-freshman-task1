package utils

import (
	"bingyan-freshman-task0/internal/config"
	"context"
	"math"
	"math/rand"
	"strconv"
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

func GenerateValidationCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}

func WriteValidationCode(email string, code string) error {
	ctx := context.Background()
	err := RedisClient.Set(ctx, email, code, 0).Err()
	if err != nil {
		return err
	}
	err = RedisClient.Expire(ctx, email, time.Duration(config.Config.Captcha.Expire)*time.Second).Err()
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

func CheckEmailExist(email string) (bool, time.Time, error) {
	ctx := context.Background()
	t, err := RedisClient.TTL(ctx, email).Result()
	if t == -2 {
		return false, time.Now(), nil
	}
	if err != nil {
		return false, time.Now(), err
	}
	if config.Config.Captcha.Resend-int(math.Round(t.Seconds())) < 0 {
		return true, time.Now().Add(t), nil
	}
	return false, time.Now(), nil
}
