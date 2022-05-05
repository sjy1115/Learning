package models

import "time"

type Course struct {
	Id           int       `json:"id" gorm:"primary_key;column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Introduction string    `json:"introduction" gorm:"column:introduction"`
	Semester     string    `json:"semester" gorm:"column:semester"`
	Avatar       string    `json:"avatar" gorm:"column:avatar"`
	InviteCode   string    `json:"invite_code" gorm:"column:invite_code"`
	InsertTm     time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm     time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (c *Course) TableName() string {
	return "course"
}
