package models

import "time"

type CourseUser struct {
	ID       int       `json:"id" gorm:"column:id;primary_key"`
	CourseID int       `json:"course_id" gorm:"column:course_id"`
	UserID   int       `json:"user_id" gorm:"column:user_id"`
	Status   int       `json:"status" gorm:"column:status"`
	InsertTm time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (cs *CourseUser) TableName() string {
	return "course_user"
}
