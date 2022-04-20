package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
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
	Username   string `yaml:"username" json:"username"`
	Password   string `yaml:"password" json:"password"`
	DB         int    `yaml:"db" json:"db"`
}

func (c *Config) Load() {
	c.Docker = false
	c.Addr = "127.0.0.1"
	c.DockerAddr = "redis"
	c.Port = 6379
	c.Password = "admin"
	c.DB = 1
}

func InitRedis(cfg *Config) {
	once.Do(func() {
		addr := fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)
		if cfg.Docker {
			addr = fmt.Sprintf("%s:%d", cfg.DockerAddr, cfg.Port)
		}

		redisCli = redis.NewClient(&redis.Options{
			Addr:     addr,
			Username: cfg.Username,
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

func RedisCmd() *redis.Client {
	return redisCli
}
