package models

type User struct {
	Id       int    `json:"id"`
	RoleId   int    `json:"role_id"`
	Age      int    `json:"age"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"`
	Avatar   string `json:"avatar"`
}

func (u User) TableName() string {
	return "users"
}
