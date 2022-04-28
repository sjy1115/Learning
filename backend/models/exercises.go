package models

import "time"

type Exercises struct {
	Id       int       `json:"id" gorm:"primary_key;column:id"`
	Title    string    `json:"title" gorm:"column:title"`
	CourseId int       `json:"course_id" gorm:"column:course_id"`
	InsertTm time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (e *Exercises) TableName() string {
	return "exercises"
}
