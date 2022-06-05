package repository

import (
	"context"
	"time"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/go-redis/redis/v9"
)

var (
	REDIS     *redis.Client
	RedisRepo *RedisImpl
)

func ConnectRedis() (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     configs.Config.REDIS_URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err = rdb.Set(context.Background(), "test", "test", time.Second*1).Err()
	return
}

type RedisImpl struct{}

func (r *RedisImpl) IsJwtValid(token string) bool {
	err := REDIS.Get(context.Background(), token).Err()
	return err != redis.Nil
}

func (r *RedisImpl) AddJwt(token string, expire time.Duration) error {
	return REDIS.Set(context.Background(), token, "", expire).Err()
}

func (r *RedisImpl) RemoveJwt(token string) error {
	return REDIS.Del(context.Background(), token).Err()
}
