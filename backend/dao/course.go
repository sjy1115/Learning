package dao

import (
	"context"
	"learning/consts"
	"learning/db/mysql"
	"learning/models"
)

func CourseGetById(ctx context.Context, id int) (course models.Course, err error) {
	err = mysql.GetRds(ctx).
		Model(&course).
		Where("id = ?", id).
		First(&course).
		Error
	return
}

func CourseGetByInviteCode(ctx context.Context, inviteCode string) (course models.Course, err error) {
	err = mysql.GetRds(ctx).
		Model(&course).
		Where("invite_code = ?", inviteCode).
		First(&course).
		Error
	return
}

func CourseUpdateById(ctx context.Context, id int, data interface{}) error {
	return mysql.GetRds(ctx).
		Model(&models.Course{}).
		Where("id = ?", id).
		UpdateColumns(data).
		Error
}

func TeacherGetById(ctx context.Context, id int) (teacher models.User, err error) {
	err = mysql.GetRds(ctx).
		Model(&teacher).
		Where("id = ? AND role = ?", id, consts.RoleTeacher).
		First(&teacher).
		Error
	return
}

func CourseDeleteById(ctx context.Context, id int) error {
	return mysql.GetRds(ctx).
		Model(&models.Course{}).
		Where("id = ?", id).
		Delete(&models.Course{}).
		Error
}

func TeacherGetByCourseId(ctx context.Context, courseId int) (user models.User, err error) {
	err = mysql.GetRds(ctx).
		Model(&user).
		Joins("LEFT JOIN course_user u ON user.id = u.user_id AND user.role = ?", consts.RoleTeacher).
		Where("u.course_id = ?", courseId).
		Error
	return
}

func StudentNumGetByCourseId(ctx context.Context, courseId int) (num int64, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.CourseUser{}).
		Joins("LEFT JOIN user u ON u.id = course_user.user_id AND u.role = ?", consts.RoleStudent).
		Where("course_user.course_id = ?", courseId).
		Count(&num).
		Error
	return
}
