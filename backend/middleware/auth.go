package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//err := jwt.Parse(c)
		//if err != nil {
		//	//TODO log
		//	c.JSON(401, gin.H{
		//		"code": 401,
		//		"msg":  err.Error(),
		//	})
		//	c.Abort()

		//}
		c.Next()
	}
}
