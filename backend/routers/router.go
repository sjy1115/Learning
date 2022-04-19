package routers

import (
	"learning/middleware"
	v1 "learning/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(router gin.IRouter) {
	router.Use(gin.Recovery(), gin.Logger(), middleware.Auth())

	api := router.Group("/api")
	{
		v1.RegisterStudentsRouter(api)
		v1.RegisterStaticRouter(api)
		v1.RegisterCourseRouter(api)
	}
}
