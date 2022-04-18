package oss

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"students/pkg/errors"
	"students/utils"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	*Config

	client *minio.Client
}

func (a *Minio) Check() (err error) {
	a.client, err = minio.New(a.endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(a.accessKey, a.accessSecret, ""),
	})
	if err != nil {
		return errs.Wrap("oss", err)
	}

	return nil
}

func (a *Minio) List(prefix string) ([]string, error) {
	objs := make([]string, 0)
	objectChan := a.client.ListObjects(context.TODO(), a.bucketName, minio.ListObjectsOptions{Prefix: prefix})
	for v := range objectChan {
		if v.Err != nil {
			break
		}

		objs = append(objs, v.Key)
	}

	return objs, nil
}

func (a *Minio) Url(key string) (string, error) {
	u, err := a.client.PresignedGetObject(context.TODO(), a.bucketName, key, time.Second*time.Duration(a.expire), nil)
	if err != nil {
		return "", err
	}

	if a.pubAddr != "" {
		u.Host = a.pubAddr
	}

	return u.String(), nil
}

func (a *Minio) Path(key string, dirs ...string) (string, error) {
	readCloser, err := a.client.GetObject(context.TODO(), a.bucketName, key, minio.GetObjectOptions{})
	if err != nil {
		return "", errs.Wrap("oss path", err)
	}
	defer readCloser.Close()

	finalDir := append([]string{a.targetPath}, dirs...)
	finalDir = append(finalDir, key)

	path := filepath.Join(finalDir...)
	if !utils.PathExist(filepath.Dir(path)) {
		err = os.MkdirAll(filepath.Dir(path), os.FileMode(0755))
		if err != nil {
			return "", errs.Wrap("oss path", err)
		}
	}

	targetF, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		return "", errs.Wrap("aliyun path", err)
	}
	defer targetF.Close()

	_, err = io.Copy(targetF, readCloser)
	if err != nil {
		return "", errs.Wrap("oss path", err)
	}

	return path, nil
}

func (a *Minio) Content(key string) ([]byte, error) {
	path, err := a.Path(key)
	if err != nil {
		return nil, errs.Wrap("oss content", err)
	}
	defer os.Remove(path)
	return ioutil.ReadFile(path)
}

func (a *Minio) Put(key, file string) error {
	_, err := a.client.FPutObject(context.TODO(), a.bucketName, key, file, minio.PutObjectOptions{})
	return err
}

func (a *Minio) PutReader(key string, reader io.Reader, size int64) error {
	_, err := a.client.PutObject(context.TODO(), a.bucketName, key, reader, size, minio.PutObjectOptions{})
	return err
}

func (a *Minio) Del(key string) error {
	return a.client.RemoveObject(context.TODO(), a.bucketName, key, minio.RemoveObjectOptions{})
}
