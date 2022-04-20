package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"learning/consts"
	"learning/pkg/jwt"
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

			c.JSON(401, gin.H{
				"code": 401,
				"msg":  err.Error(),
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
