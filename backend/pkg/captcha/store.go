package captcha

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

type RedisStore struct {
	cli *redis.Client
}

//Set(id string, value string) error
//    Get(id string, clear bool) string
//    Verify(id string, answer string, clear bool) bool

func newRedisStore(cli *redis.Client) *RedisStore {
	return &RedisStore{
		cli: cli,
	}
}

func (r *RedisStore) Set(id, value string) error {
	ctx := context.Background()

	_, err := r.cli.Set(ctx, id, value, time.Second*60).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	resp, err := r.cli.Get(ctx, id).Result()
	if err != nil {
		return ""
	}

	if clear {
		r.cli.Del(ctx, id).Result()
	}

	return resp
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	vv := r.Get(id, clear)

	vv = strings.TrimSpace(vv)
	return vv == strings.TrimSpace(answer)
}
