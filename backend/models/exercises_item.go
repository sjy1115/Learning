package models

import "time"

type ExerciseItem struct {
	Id         int       `json:"id" gorm:"primary_key;column:id"`
	ExerciseId int       `json:"exercise_id" gorm:"column:exercise_id"`
	Title      string    `json:"title" gorm:"column:title"`
	Type       string    `json:"type" gorm:"column:type"`
	Answer     string    `json:"answer" gorm:"column:answer"`
	Options    string    `json:"options" gorm:"column:options"`
	InsertTm   time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm   time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (q *ExerciseItem) TableName() string {
	return "exercise_item"
}
