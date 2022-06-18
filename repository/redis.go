package repository

import (
	"context"
	"time"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/go-redis/redis/v9"
)

var (
	RedisDefault  *defaultImpl
	RedisJwt      *jwtImpl
	RedisCallback *callbackImpl
)

type RedisDatabase int

const (
	DEFAULT  RedisDatabase = 0
	JWT      RedisDatabase = 1
	CALLBACK RedisDatabase = 2
)

func ConnectRedis(db RedisDatabase) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     configs.Config.REDIS_URL,
		Password: "",      // no password set
		DB:       int(db), // use default DB
	})
	switch db {
	case JWT:
		RedisJwt = &jwtImpl{Client: rdb}
	case CALLBACK:
		RedisCallback = &callbackImpl{Client: rdb}
	default:
		RedisDefault = &defaultImpl{Client: rdb}
	}

	err = rdb.Set(context.Background(), "test", "test", time.Second*1).Err()
	return
}

func IsRedisReady() error {
	err := RedisDefault.Client.Set(context.Background(), "test", "test", time.Second*1).Err()
	if err != nil {
		return err
	}
	err = RedisJwt.Client.Set(context.Background(), "test", "test", time.Second*1).Err()
	if err != nil {
		return err
	}
	return RedisCallback.Client.Set(context.Background(), "test", "test", time.Second*1).Err()
}

type defaultImpl struct {
	Client *redis.Client
}

type jwtImpl struct {
	Client *redis.Client
}

type callbackImpl struct {
	Client *redis.Client
}

func (r *jwtImpl) IsJwtValid(token string) bool {
	err := r.Client.Get(context.Background(), token).Err()
	return err != redis.Nil
}

func (r *jwtImpl) AddJwt(token string, expire time.Duration) error {
	return r.Client.Set(context.Background(), token, "", expire).Err()
}

func (r *jwtImpl) RemoveJwt(token string) error {
	return r.Client.Del(context.Background(), token).Err()
}

func (r *callbackImpl) AddCallback(state, callback string) error {
	return r.Client.Set(context.Background(), state, callback, time.Minute*15).Err()
}

func (r *callbackImpl) GetCallback(state string) (string, error) {
	return r.Client.Get(context.Background(), state).Result()
}
