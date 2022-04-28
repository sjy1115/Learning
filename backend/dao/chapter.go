package dao

import (
	"context"
	"gorm.io/gorm"
	"learning/db/mysql"
	"learning/models"
)

func ChapterGetById(ctx context.Context, id int) (chapter models.Chapter, err error) {
	err = mysql.GetRds(ctx).
		Model(&chapter).
		Where("id = ?", id).
		First(&chapter).
		Error
	return
}

func ChapterUpdateById(ctx context.Context, chapterId int, data interface{}, tx ...*gorm.DB) (err error) {
	var db *gorm.DB
	if len(tx) != 0 {
		db = tx[0]
	} else {
		db = mysql.GetRds(ctx)
	}
	err = db.Model(&models.Chapter{}).
		Where("chapter_id = ?", chapterId).
		UpdateColumns(data).
		Error
	return
}

func ChapterDeleteById(ctx context.Context, chapterId int) error {
	return mysql.GetRds(ctx).
		Model(&models.Chapter{}).
		Where("id = ?", chapterId).
		Delete(&models.Chapter{}).
		Error
}
