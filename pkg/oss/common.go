package oss

import (
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"strconv"
)

type Type int

const (
	_ Type = iota
	TypeAliyun
	TypeFileServer
	TypeMinio
)

var Bucket Oss

type Oss interface {
	List(prefix string) ([]string, error)
	Url(key string) (string, error)
	Path(key string, dir ...string) (string, error)
	Content(key string) ([]byte, error)
	Put(key, file string) error
	PutReader(key string, reader io.Reader, size int64) error
	Del(key string) error
}

func New(chs ...ConfigHandler) Oss {
	cfg := &Config{}

	for _, ch := range chs {
		ch(cfg)
	}

	err := cfg.Check()
	if err != nil {
		panic(err)
	}

	var o Oss
	switch cfg.t {
	case TypeMinio:
		o = &Minio{Config: cfg}
	default:
		panic("unsupported oss type")
	}

	checker, ok := o.(interface {
		Check() error
	})
	if ok {
		err = checker.Check()
		if err != nil {
			log.Fatal("check oss", zap.Error(err), zap.Any("config", cfg))
		}
	}

	return o
}

func init() {
	ossType := os.Getenv("moscan_oss_type")
	if ossType == "" {
		ossType = "3"
	}

	intType, err := strconv.Atoi(ossType)
	if err != nil {
		panic(err)
	}

	Bucket = New(
		WithType(Type(intType)),
		WithPrefix("/tmp/moresec/"),
		WithAccessKey(os.Getenv("moscan_oss_accessKey")),
		WithAccessSecret(os.Getenv("moscan_oss_accessSecret")),
		WithEndpoint(os.Getenv("moscan_oss_endpoint")),
		WithBucketName(os.Getenv("moscan_oss_bucket")),
		WithExpire(60*15),
		WithPubAddr(os.Getenv("moscan_oss_pub_addr")),
	)
}
