package main

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/sirupsen/logrus"
)

type RedisManager struct {
	Host     string
	Port     string
	DB       int
	Password string
	Client   *redis.Client
}

func NewRedisManager(host string, port string, db int, password string) (*RedisManager, error) {
	redisMgr := &RedisManager{
		Host:     host,
		Port:     port,
		DB:       db,
		Password: password,
	}
	redisMgr.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})
	_, err := redisMgr.Client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisMgr, nil
}

func (redisMgr *RedisManager) SetStringValue(key, value string) error {
	return redisMgr.Client.Set(key, value, 0).Err()
}

func (redisMgr *RedisManager) GetStringValue(key string) (string, error) {
	return redisMgr.Client.Get(key).Result()
}
