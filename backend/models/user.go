package models

import "time"

type User struct {
	Id       int       `json:"id" gorm:"primary_key;column:id"`
	Name     string    `json:"name" gorm:"column:name"`
	Phone    string    `json:"phone" gorm:"column:phone"`
	College  string    `json:"college" gorm:"column:college"`
	Role     int       `json:"role" gorm:"column:role"`
	Gender   string    `json:"gender" gorm:"column:gender"`
	Number   string    `json:"number" gorm:"column:number"`
	Password string    `json:"password" gorm:"column:password"`
	Avatar   string    `json:"avatar" gorm:"column:avatar"`
	InsertTm time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (u User) TableName() string {
	return "user"
}
