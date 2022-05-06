package proto

type LoginRequest struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	CaptId   string `json:"capt_id" form:"capt_id"`
	Vcode    string `json:"vcode" form:"vcode"`
}

type LoginResponse struct {
	UserId int    `json:"user_id"`
	Role   int    `json:"role"`
	Token  string `json:"token"`
}

type UserUpdateRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type UserUpdateResponse struct {
}

type LogoutResponse struct {
}

type UserInfoResponse struct {
	UserId   int    `json:"user_id"`
	Role     int    `json:"role"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
	College  string `json:"college" form:"college"`
	Role     int    `json:"role" form:"role"`
	Gender   string `json:"gender" form:"gender"`
	Number   string `json:"number" form:"number"`
	Password string `json:"password" form:"password"`
}

type RegisterResponse struct {
	Phone    string `json:"phone" form:"phone"`
	School   string `json:"school" form:"school"`
	ROle     string `json:"role" form:"role"`
	Gender   int    `json:"gender" form:"gender"`
	Number   string `json:"number" form:"number"`
	Password string `json:"password" form:"password"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
}

type ChangePasswordResponse struct {
}

type VerifyCodeResponse struct {
	CaptId string `json:"capt_id" form:"capt_id"`
	Image  string `json:"image" form:"image"`
}
