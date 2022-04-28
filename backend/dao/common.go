package dao

import (
	"context"
	"learning/db/mysql"

	"gorm.io/gorm"
)

func GetById(ctx context.Context, id int, result interface{}) error {
	return mysql.GetRds(ctx).Model(result).Where("id = ?", id).First(result).Error
}

func Create(ctx context.Context, data interface{}, tx ...*gorm.DB) error {
	mdb := mysql.GetRds(ctx)
	if len(tx) > 0 {
		mdb = tx[0]
	}
	return mdb.Model(data).Create(data).Error
}

func UpdateColumn(tx *gorm.DB, id int, data interface{}) error {
	return tx.Model(data).Where("id = ?", id).UpdateColumns(data).Error
}
