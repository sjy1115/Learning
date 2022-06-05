package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"learning/consts"
	"learning/db/cache"
	"learning/pkg/jwt"
	"net/http"
	"strings"
)

var (
	ignoreUrls []string
)

func InitIgnoreUrl(urls []string) {
	ignoreUrls = urls
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if checkUrl(c.Request.URL.Path) || strings.HasPrefix(path, "/api/static") {
			c.Next()
			return
		}

		token := c.GetHeader(consts.AuthHeader)

		err := jwt.Parse(c)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"token": token,
				"error": err.Error(),
			}).Error("jwt auth failed")

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
			})
			c.Abort()
		}

		exists, _ := cache.Exist(c.Request.Context(), cache.UserTokenKey(c.MustGet(consts.AuthToken).(*jwt.UserToken).UserId))
		if !exists {
			logrus.WithFields(logrus.Fields{
				"token":   token,
				"user id": c.MustGet(consts.AuthToken).(*jwt.UserToken).UserId,
			}).Error("user not exists")

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "user not exists",
			})
			c.Abort()
		}

		rdsToken, _ := cache.Get(c.Request.Context(), cache.UserTokenKey(c.MustGet(consts.AuthToken).(*jwt.UserToken).UserId))
		if token != rdsToken {
			logrus.WithFields(logrus.Fields{
				"token":     token,
				"rds token": rdsToken,
				"user id":   c.MustGet(consts.AuthToken).(*jwt.UserToken).UserId,
			}).Error("token not match")

			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "token not match",
			})
			c.Abort()
		}

		c.Next()
	}
}

func checkUrl(path string) bool {
	for _, v := range ignoreUrls {
		if (v[len(v)-1] == '*' && strings.HasPrefix(path, v[:len(v)-1])) || path == v {
			return true
		}
	}
	return false
}
