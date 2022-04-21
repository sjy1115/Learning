package captcha

import (
	"context"
	"github.com/sirupsen/logrus"
	"learning/db/cache"
	"strings"
)

type RedisStore struct {
}

//Set(id string, value string) error
//    Get(id string, clear bool) string
//    Verify(id string, answer string, clear bool) bool

func (r *RedisStore) Set(id, value string) error {
	ctx := context.Background()

	err := cache.SetEx(ctx, id, value, 60)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id":  id,
			"err": err,
		}).Error("set captcha")
		return err
	}

	return nil
}

func (r *RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	resp, err := cache.Get(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id": id,
		}).Error("get captcha")
		return ""
	}

	if clear {
		err = cache.Del(ctx, id)
		logrus.WithFields(logrus.Fields{
			"id": id,
		}).Error("delete captcha")
	}

	return resp
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	vv := r.Get(id, clear)

	vv = strings.TrimSpace(vv)
	return vv == strings.TrimSpace(answer)
}
