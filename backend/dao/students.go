package dao

import (
	"context"
	"learning/db/mysql"

	"learning/models"
)

func CreateUser(ctx context.Context, student *models.User) error {
	return Create(mysql.GetRds(ctx), student)
}

func GetAllUsers(ctx context.Context) (results []models.User, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.User{}).
		Find(&results).
		Error
	return
}

func GetPageUser(ctx context.Context, page, pageSize int) (results []models.User, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.User{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&results).
		Error
	return
}

func GetUserById(ctx context.Context, id int) (result models.User, err error) {
	err = GetById(ctx, id, &result)
	return
}

func UpdateUserById(ctx context.Context, id int) error {
	return nil
}

func DeleteUserById(ctx context.Context, id int) error {
	return mysql.GetRds(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Delete(&models.User{}).
		Error
}
