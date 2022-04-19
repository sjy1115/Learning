package dao

import (
	"context"
	"learning/db"

	"gorm.io/gorm"
)

func GetById(ctx context.Context, id int, result interface{}) error {
	return db.GetRds(ctx).Model(result).Where("id = ?", id).First(result).Error
}

func Create(tx *gorm.DB, data interface{}) error {
	return tx.Model(data).Create(data).Error
}
