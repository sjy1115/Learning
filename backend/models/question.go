package models

import "time"

type Question struct {
	Id       int       `json:"id" gorm:"primary_key;column:id"`
	Topic    string    `json:"topic" gorm:"column:topic"`
	Type     string    `json:"type" gorm:"column:type"`
	Answer   string    `json:"answer" gorm:"column:answer"`
	Options  string    `json:"options" gorm:"column:options"`
	InsertTm time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (q *Question) TableName() string {
	return "question"
}
