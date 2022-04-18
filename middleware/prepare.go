package middleware

import "github.com/gin-gonic/gin"

func PrepareMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if len(path) > 1 && path[len(path)-1] == '/' && path[len(path)-2] != '/' {
			path = path[:len(path)-1]
			c.Request.RequestURI = path
			c.Request.URL.Path = path
		}
	}
}
