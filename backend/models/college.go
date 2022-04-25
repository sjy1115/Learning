package models

import "time"

type College struct {
	Id       int       `json:"id" gorm:"primary_key;column:id"`
	Name     string    `json:"name" gorm:"column:name"`
	InsertTm time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (c *College) TableName() string {
	return "college"
}
