package jwt

import (
	sjwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"learning/config"
	"learning/consts"
	"time"
)

type UserToken struct {
	UserId   int    `json:"user_id"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	sjwt.StandardClaims
}

func GenerateToken(uerId, role int, username, password, phone string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24)

	claims := &UserToken{
		UserId:   uerId,
		Username: username,
		Password: password,
		Phone:    phone,
		Role:     role,
		StandardClaims: sjwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-learning",
		},
	}

	tokenClaims := sjwt.NewWithClaims(sjwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(config.Conf.Jwt.Secret))
	return token, err
}

func Parse(c *gin.Context) error {
	token := c.GetHeader(consts.AuthHeader)

	userToken, err := ParseToken(token)
	if err != nil {
		return err
	}

	c.Set(consts.AuthToken, userToken)

	return nil
}

func ParseToken(token string) (*UserToken, error) {
	tokenClaims, err := sjwt.ParseWithClaims(token, &UserToken{}, func(token *sjwt.Token) (interface{}, error) {
		return []byte(config.Conf.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*UserToken); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, nil
}

func GetToken(c *gin.Context) *UserToken {
	token, _ := c.Get(consts.AuthToken)
	if token == nil {
		return nil
	}
	return token.(*UserToken)
}
