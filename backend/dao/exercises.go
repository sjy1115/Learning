package dao

import (
	"context"
	"learning/db/mysql"
	"learning/models"
)

func ExercisesGetById(ctx context.Context, chapterId int) (exercises models.Exercises, err error) {
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
