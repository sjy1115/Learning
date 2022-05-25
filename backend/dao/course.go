package dao

import (
	"context"
	"learning/consts"
	"learning/db/mysql"
	"learning/models"
	"time"
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
		First(&user).
		Error
	return
}

func StudentNumGetByCourseId(ctx context.Context, courseId int) (num int64, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.CourseUser{}).
		Joins("JOIN user u ON u.id = course_user.user_id AND u.role = ?", consts.RoleStudent).
		Where("course_user.course_id = ?", courseId).
		Count(&num).
		Error
	return
}

func StudentLearnNumBetByChapterId(ctx context.Context, chapterId int) (num int64, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.ChapterUser{}).
		Joins("JOIN user u ON u.id = chapter_user.user_id AND u.role = ?", consts.RoleStudent).
		Where("chapter_user.chapter_id = ?", chapterId).
		Count(&num).
		Error

	return
}

func StudentIsLearnedByStudentId(ctx context.Context, userId, chapterId int) (cus []models.ChapterUser, err error) {
	err = mysql.GetRds(ctx).
		Model(&models.ChapterUser{}).
		Where("user_id = ? AND chapter_id = ?", userId, chapterId).
		Find(&cus).
		Error

	return
}

func ScoreGetByChapterIdAndUserId(ctx context.Context, userId, chapterId int) (score int, exists bool, err error) {
	var cus []models.ChapterUser

	err = mysql.GetRds(ctx).
		Model(&models.ChapterUser{}).
		Where("user_id = ? AND chapter_id = ?", userId, chapterId).
		Find(&cus).
		Error

	if len(cus) > 0 {
		score = cus[0].Score
		exists = true
	}

	return
}

func UpdateStudentScoreByStudentId(ctx context.Context, userId, chapterId, score int) error {
	oldScore, exists, _ := ScoreGetByChapterIdAndUserId(ctx, userId, chapterId)

	if !exists {
		return mysql.GetRds(ctx).Create(&models.ChapterUser{
			ChapterID: chapterId,
			UserID:    userId,
			Score:     score,
			InsertTm:  time.Now(),
			UpdateTm:  time.Now(),
		}).Error
	} else {
		if score > oldScore {
			return mysql.GetRds(ctx).
				Model(&models.ChapterUser{}).
				Where("user_id = ? AND chapter_id = ?", userId, chapterId).
				Updates(map[string]interface{}{
					"score":     score,
					"update_tm": time.Now(),
				}).Error
		}
		return nil
	}
}
