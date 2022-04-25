package models

import "time"

type Exam struct {
	Id       int       `json:"id" gorm:"primary_key;column:id"`
	Name     string    `json:"name" gorm:"column:name"`
	CourseId int       `json:"course_id" gorm:"column:course_id"`
	InsertTm time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (e *Exam) TableName() string {
	return "exam"
}
