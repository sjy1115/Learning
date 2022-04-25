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
