package jwt

import (
	sjwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"learning/consts"
	"time"
)

var jwtSecret []byte

type UserToken struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	sjwt.StandardClaims
}

func GenerateToken(uerId, role int, username, password string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24)

	userToken := &UserToken{
		UserId:   uerId,
		Username: username,
		Password: password,
		Role:     role,
		StandardClaims: sjwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-learning",
		},
	}

	token := sjwt.NewWithClaims(sjwt.SigningMethodES256, userToken)

	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(c *gin.Context) error {
	token := c.GetHeader(consts.AuthHeader)

	tokenClaims, err := sjwt.ParseWithClaims(token, &UserToken{}, func(token *sjwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := tokenClaims.Claims.(*UserToken); ok && tokenClaims.Valid {
		c.Set(consts.AuthToken, claims)
		return nil
	}

	return nil
}

func GetToken(c *gin.Context) *UserToken {
	token, _ := c.Get(consts.AuthToken)
	return token.(*UserToken)
}

func InitJwtToken(secret string) {
	jwtSecret = []byte(secret)
}
