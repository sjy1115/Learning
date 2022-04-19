package models

type User struct {
	Id       int    `json:"id"`
	Age      int    `json:"age"`
	Name     string `json:"name"`
	Password string
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"`
	Photo    string `json:"photo"`
}

func (u User) TableName() string {
	return "users"
}
