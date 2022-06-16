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
		Updates(data).
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

func ChapterUserGetByChapterIdAndUserId(ctx context.Context, userId, chapterId int) (cu models.ChapterUser, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.ChapterUser{}).
		Where("user_id = ?", userId).
		Where("chapter_id = ?", chapterId).
		First(&cu).
		Error

	return
}

func ChapterUserUpdateByChapterIdAndUserId(ctx context.Context, userId, chapterId int, data map[string]interface{}) error {
	return mysql.GetRds(ctx).
		Model(&models.ChapterUser{}).
		Where("user_id = ?", userId).
		Where("chapter_id = ?", chapterId).
		Updates(data).
		Error
}

func UserSignInStatus(ctx context.Context, userId, chapterId int) (signIn bool, err error) {
	var cu models.ChapterUser

	err = mysql.GetRds(ctx).Model(&cu).Where("user_id = ? AND chapter_id = ?", userId, chapterId).First(&cu).Error
	if err != nil {
		return
	}

	return cu.SignIn == 1 && cu.Status == 1, nil
}
