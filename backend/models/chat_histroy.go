package models

import "time"

type ChatHistory struct {
	Id       int       `json:"id" gorm:"primary_key;column:id"`
	RoomId   int       `json:"room_id" gorm:"column:room_id"`
	From     int       `json:"from" gorm:"column:from"`
	To       int       `json:"to" gorm:"column:to"`
	Text     string    `json:"text" gorm:"column:text"`
	InsertTm time.Time `json:"insert_tm" gorm:"column:insert_tm"`
}

func (c *ChatHistory) TableName() string {
	return "chat_history"
}
