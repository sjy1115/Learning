package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"learning/dao"
	"learning/db/cache"
	"learning/models"
	"learning/pkg/captcha"
	"learning/pkg/context"
	"learning/pkg/jwt"
	"learning/proto"
	"strconv"
)

func UserLoginHandler(c *context.Context, req *proto.LoginRequest) (resp *proto.LoginResponse, err error) {
	if !captcha.VeryCaptcha(req.CaptId, req.Vcode) {
		return nil, fmt.Errorf("验证码错误")
	}

	user, err := dao.GetUserByPhone(c.Ctx, req.Phone)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("密码错误")
	}

	token, err := jwt.GenerateToken(user.Id, user.Role, user.Name, "", req.Phone)
	if err != nil {
		return nil, err
	}

	// token 有限期为 1 小时
	err = cache.SetEx(c.Ctx, cache.UserTokenKey(user.Id), token, 3600)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		UserId: user.Id,
		Role:   user.Role,
		Token:  token,
	}, nil
}

func UserInfoHandler(c *context.Context) (resp *proto.UserInfoResponse, err error) {
	user, err := dao.GetUserById(c.Ctx, c.UserToken.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.UserInfoResponse{
		UserId:   user.Id,
		Role:     user.Role,
		Username: user.Name,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
	}, nil
}

func UserUpdateHandler(c *context.Context, req *proto.UserUpdateRequest) (resp *proto.UserUpdateResponse, err error) {
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)

	user, err := dao.GetUserById(c.Ctx, userId)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	if len(req.Avatar) != 0 {
		data["avatar"] = req.Avatar
	}

	if len(req.Name) != 0 {
		data["name"] = req.Name
	}

	err = dao.UserUpdateById(c.Ctx, user.Id, data)
	if err != nil {
		return nil, err
	}

	return
}

func UserLogoutHandler(c *context.Context) (resp *proto.LogoutResponse, err error) {
	err = cache.Del(c.Ctx, cache.UserTokenKey(c.UserToken.UserId))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func UserRegisterHandler(c *context.Context, req *proto.RegisterRequest) (resp *proto.RegisterResponse, err error) {
	exists, err := dao.UserExistsByName(c.Ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("用户名已存在")
	}

	exists, err = dao.UserExistsByPhone(c.Ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("手机号已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Role:     req.Role,
		Name:     req.Username,
		College:  req.College,
		Gender:   req.Gender,
		Number:   req.Number,
		Phone:    req.Phone,
		Password: string(hashedPassword),
	}

	err = dao.CreateUser(c.Ctx, &user)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func VerifyCodeHandler(c *context.Context) (resp *proto.VerifyCodeResponse, err error) {
	resp = new(proto.VerifyCodeResponse)

	id, image, err := captcha.GenerateCaptcha()
	if err != nil {
		return nil, err
	}

	resp.CaptId = id
	resp.Image = image

	return
}
