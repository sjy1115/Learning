package oss

import (
	"path/filepath"
	"students/pkg/errors"
)

type Config struct {
	accessKey    string
	accessSecret string
	endpoint     string
	pubAddr      string
	bucketName   string
	prefix       string
	t            Type
	expire       int64

	targetPath string
}

func (c *Config) Check() error {
	if c.accessKey == "" {
		return errs.Wrap("access_key", errs.Empty)
	}
	if c.accessSecret == "" {
		return errs.Wrap("access_secret", errs.Empty)
	}
	if c.endpoint == "" {
		return errs.Wrap("endpoint", errs.Empty)
	}
	if c.bucketName == "" {
		return errs.Wrap("bucketName", errs.Empty)
	}

	if c.targetPath == "" {
		c.targetPath = c.prefix
	}
	c.targetPath = filepath.Clean(c.targetPath)

	return nil
}

type ConfigHandler func(c *Config)

func WithAccessKey(ak string) ConfigHandler {
	return func(c *Config) {
		c.accessKey = ak
	}
}

func WithAccessSecret(as string) ConfigHandler {
	return func(c *Config) {
		c.accessSecret = as
	}
}

func WithEndpoint(endpoint string) ConfigHandler {
	return func(c *Config) {
		c.endpoint = endpoint
	}
}

func WithBucketName(bucketName string) ConfigHandler {
	return func(c *Config) {
		c.bucketName = bucketName
	}
}

func WithPrefix(prefix string) ConfigHandler {
	return func(c *Config) {
		c.prefix = prefix
	}
}

func WithType(t Type) ConfigHandler {
	return func(c *Config) {
		c.t = t
	}
}

func WithExpire(e int64) ConfigHandler {
	return func(c *Config) {
		c.expire = e
	}
}

func WithTargetPath(tp string) ConfigHandler {
	return func(c *Config) {
		c.targetPath = tp
	}
}

func WithPubAddr(pa string) ConfigHandler {
	return func(c *Config) {
		c.pubAddr = pa
	}
}
