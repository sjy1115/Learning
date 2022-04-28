package models

import "time"

type ChapterUser struct {
	ID        int       `json:"id" gorm:"primary_key;column:id"`
	UserID    int       `json:"user_id" gorm:"column:user_id"`
	ChapterID int       `json:"chapter_id" gorm:"column:chapter_id"`
	Status    int       `json:"status" gorm:"column:status"`
	InsertTm  time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm  time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (cu *ChapterUser) TableName() string {
	return "chapter_user"
}
