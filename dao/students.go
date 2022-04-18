package dao

import (
	"context"

	"students/db"
	"students/models"
)

func CreateStudent(ctx context.Context, student *models.Student) error {
	return Create(db.GetRds(ctx), student)
}

func GetAllStudents(ctx context.Context) (results []models.Student, err error) {
	err = db.GetRds(ctx).
		Model(&models.Student{}).
		Find(&results).
		Error
	return
}

func GetPageStudents(ctx context.Context, page, pageSize int) (results []models.Student, err error) {
	err = db.GetRds(ctx).
		Model(&models.Student{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&results).
		Error
	return
}

func GetStudentById(ctx context.Context, id int) (result models.Student, err error) {
	err = GetById(ctx, id, &result)
	return
}

func UpdateStudentById(ctx context.Context, id int) error {
	return nil
}

func DeleteStudentById(ctx context.Context, id int) error {
	return db.GetRds(ctx).
		Model(&models.Student{}).
		Where("id = ?", id).
		Delete(&models.Student{}).
		Error
}
