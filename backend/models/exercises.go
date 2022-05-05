package models

import "time"

type Exercises struct {
	Id         int       `json:"id" gorm:"primary_key;column:id"`
	Title      string    `json:"title" gorm:"column:title"`
	ChapterId  int       `json:"chapter_id" gorm:"column:chapter_id"`
	Attachment string    `json:"attachment" gorm:"column:attachment"`
	InsertTm   time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm   time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (e *Exercises) TableName() string {
	return "exercises"
}
