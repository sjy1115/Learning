package models

type User struct {
	Id       int    `json:"id" gorm:"primary_key;column:id"`
	Username string `json:"username" gorm:"column:username"`
	Phone    string `json:"phone" gorm:"column:phone"`
	College  string `json:"college" gorm:"column:college"`
	Role     int    `json:"role" gorm:"column:role"`
	Gender   string `json:"gender" gorm:"column:gender"`
	Number   string `json:"number" gorm:"column:number"`
	Password string `json:"password" gorm:"column:password"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
}

func (u User) TableName() string {
	return "users"
}
