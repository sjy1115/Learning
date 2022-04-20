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
	"learning/utils"
)

func UserLoginHandler(c *context.Context, req *proto.LoginRequest) (resp *proto.LoginResponse, err error) {
	if captcha.VeryCaptcha(req.CaptId, req.Vcode) {
		utils.Error(c, utils.ErrorCode, fmt.Errorf("验证码错误"))
		return
	}

	user, err := dao.GetUserByName(c.Ctx, req.Username)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.Error(c, utils.ErrorCode, fmt.Errorf("密码错误"))
		return
	}

	token, err := jwt.GenerateToken(user.Id, user.Role, user.Username, "")
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	// token 有限期为 1 小时
	err = cache.SetEx(c.Ctx, cache.UserTokenKey(user.Id), token, 3600)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	return &proto.LoginResponse{
		UserId: user.Id,
		Token:  token,
	}, nil
}

func UserLogoutHandler(c *context.Context) (resp *proto.LogoutResponse, err error) {
	// TODO: 清除缓存
	return nil, nil
}

func UserRegisterHandler(c *context.Context, req *proto.RegisterRequest) (resp *proto.RegisterResponse, err error) {
	exists, err := dao.UserExistsByName(c.Ctx, req.Username)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}
	if exists {
		utils.Error(c, utils.ErrorCode, fmt.Errorf("用户名已存在"))
	}

	exists, err = dao.UserExistsByPhone(c.Ctx, req.Phone)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}
	if exists {
		utils.Error(c, utils.ErrorCode, fmt.Errorf("手机号已存在"))
	}

	// TODO role get by id

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	user := models.User{
		Role:     0,
		Username: req.Username,
		College:  req.College,
		Gender:   req.Gender,
		Number:   req.Number,
		Password: string(hashedPassword),
	}

	err = dao.CreateUser(c.Ctx, &user)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

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
