package controllers

import (
	"learning/pkg/context"
	"learning/proto"
	"learning/services"
	"learning/utils"
)

func UserLogin(c *context.Context) {
	var req proto.LoginRequest

	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	resp, err := services.UserLoginHandler(c, &req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}

func UserLogout(c *context.Context) {
	resp, err := services.UserLogoutHandler(c)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}

func UserRegister(c *context.Context) {
	var req proto.RegisterRequest

	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	resp, err := services.UserRegisterHandler(c, &req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}

func VerifyCode(c *context.Context) {
	resp, err := services.VerifyCodeHandler(c)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}
