package proto

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	CaptId   string `json:"capt_id" form:"capt_id"`
	Vcode    string `json:"vcode" form:"vcode"`
}

type LoginResponse struct {
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}

type LogoutResponse struct {
}

type RegisterRequest struct {
}

type RegisterResponse struct {
}

type VerifyCodeResponse struct {
	CaptId string `json:"capt_id" form:"capt_id"`
	Image  string `json:"image" form:"image"`
}
