package services

import (
	"learning/pkg/captcha"
	"learning/pkg/context"
	"learning/proto"
)

func UserLoginHandler(c *context.Context, req *proto.LoginRequest) (resp *proto.LoginResponse, err error) {

	return nil, nil
}

func UserLogoutHandler(c *context.Context) (resp *proto.LogoutResponse, err error) {

	return nil, nil
}

func UserRegisterHandler(ctx *context.Context, req *proto.RegisterRequest) (resp *proto.RegisterResponse, err error) {

	return nil, nil
}

func VerifyCodeHandler(c *context.Context) (resp *proto.VerifyCodeResponse, err error) {
	id, image, err := captcha.GenerateCaptcha()
	if err != nil {
		return nil, err
	}

	resp.CaptId = id
	resp.Image = image

	return
}
