package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	once     sync.Once
	redisCli *redis.Client
)

type Config struct {
	Docker     bool   `yaml:"docker" json:"docker"`
	Addr       string `yaml:"addr" json:"addr"`
	DockerAddr string `yaml:"docker_addr" json:"docker_addr"`
	Port       int    `yaml:"port" json:"port"`
	Password   string `yaml:"password" json:"password"`
	DB         int    `yaml:"db" json:"db"`
}

func (c *Config) Load() {
	c.Docker = true
	c.Addr = "127.0.0.1"
	c.DockerAddr = "redis"
	c.Port = 6379
	c.Password = "123456"
	c.DB = 0
}

func InitRedis(cfg *Config) {
	once.Do(func() {
		addr := fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)
		if cfg.Docker {
			addr = fmt.Sprintf("%s:%d", cfg.DockerAddr, cfg.Port)
		}

		redisCli = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		})

		resp, err := redisCli.Ping(context.TODO()).Result()
		if err == nil && resp == "PONG" {
			return
		}

		panic("redis init failed")
	})
}

func Set(ctx context.Context, key string, value interface{}) error {
	return redisCli.Set(ctx, key, value, 0).Err()
}

func SetEx(ctx context.Context, key string, value string, second int) error {
	return redisCli.Set(ctx, key, value, time.Duration(second)*time.Second).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	return redisCli.Get(ctx, key).Result()
}

func Del(ctx context.Context, key string) error {
	return redisCli.Del(ctx, key).Err()
}

func Exist(ctx context.Context, key string) (bool, error) {
	num, err := redisCli.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return num > 0, nil
}

func RedisCmd() *redis.Client {
	return redisCli
}
