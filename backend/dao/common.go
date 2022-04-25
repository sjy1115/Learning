package dao

import (
	"context"
	"learning/db/mysql"

	"gorm.io/gorm"
)

func GetById(ctx context.Context, id int, result interface{}) error {
	return mysql.GetRds(ctx).Model(result).Where("id = ?", id).First(result).Error
}

func Create(tx *gorm.DB, data interface{}) error {
	return tx.Model(data).Create(data).Error
}

func UpdateColumn(tx *gorm.DB, id int, data interface{}) error {
	return tx.Model(data).Where("id = ?", id).UpdateColumns(data).Error
}
