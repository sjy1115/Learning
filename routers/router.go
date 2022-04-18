package routers

import (
	v1 "students/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(router gin.IRouter) {
	router.Use(gin.Recovery(), gin.Logger())

	api := router.Group("/api")
	{
		v1.RegisterStudentsRouter(api)
	}
}
