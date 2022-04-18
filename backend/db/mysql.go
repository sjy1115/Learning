package db

import (
	"context"
	"students/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitMysql(cfg *config.Config) error {
	dsn := cfg.Mysql.DSN()

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func GetRds(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}