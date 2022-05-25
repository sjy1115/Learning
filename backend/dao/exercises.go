package dao

import (
	"context"
	"gorm.io/gorm"
	"learning/db/mysql"
	"learning/models"
)

func ExercisesGetById(ctx context.Context, id int) (exercises models.Exercises, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.Exercises{}).
		Where("id = ?", id).
		First(&exercises).
		Error

	return
}

func ExercisesGetByChapterId(ctx context.Context, chapterId int) (exercises models.Exercises, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.Exercises{}).
		Where("chapter_id = ?", chapterId).
		First(&exercises).
		Error

	return
}

func ExerciseItemGetByExerciseId(ctx context.Context, exerciseId int) (exerciseItem []models.ExerciseItem, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.ExerciseItem{}).
		Where("exercise_id = ?", exerciseId).
		Find(&exerciseItem).
		Error

	return
}

func ExercisesUpdateById(ctx context.Context, id int, data map[string]interface{}, dbs ...*gorm.DB) (err error) {
	mdb := mysql.GetRds(ctx)
	if len(dbs) > 0 {
		mdb = dbs[0]
	}

	err = mdb.Model(&models.Exercises{}).
		Where("id = ?", id).
		Updates(data).
		Error
	return
}

func ExercisesItemDeleteByExerciseId(ctx context.Context, exerciseId int, dbs ...*gorm.DB) (err error) {
	mdb := mysql.GetRds(ctx)
	if len(dbs) > 0 {
		mdb = dbs[0]
	}

	err = mdb.Model(&models.ExerciseItem{}).
		Where("exercise_id = ?", exerciseId).
		Delete(&models.ExerciseItem{}).
		Error
	return
}
