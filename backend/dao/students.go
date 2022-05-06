package dao

import (
	"context"
	"learning/db/mysql"

	"learning/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	return Create(ctx, user)
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

func UserGetById(ctx context.Context, id int) (result models.User, err error) {
	err = GetById(ctx, id, &result)
	return
}

func UserUpdateById(ctx context.Context, id int, data map[string]interface{}) (err error) {
	err = mysql.GetRds(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		UpdateColumns(data).
		Error
	return
}

func UserExistsByName(ctx context.Context, name string) (exists bool, err error) {
	var user []models.User
	err = mysql.GetRds(ctx).
		Model(&models.User{}).
		Where("name = ?", name).
		Find(&user).
		Error
	if err != nil {
		return
	}

	return len(user) > 0, nil
}

func UserExistsByPhone(ctx context.Context, phone string) (exists bool, err error) {
	var user []models.User
	err = mysql.GetRds(ctx).
		Model(&models.User{}).
		Where("phone = ?", phone).
		Find(&user).
		Error
	if err != nil {
		return
	}

	return len(user) > 0, nil
}

func GetUserByPhone(ctx context.Context, phone string) (user models.User, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.User{}).
		Where("phone = ?", phone).
		First(&user).
		Error
	return
}

func GetUserByName(ctx context.Context, name string) (result models.User, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.User{}).
		Where("username = ?", name).
		First(&result).
		Error
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
