package v1

import (
	"github.com/gin-gonic/gin"
	"learning/controllers"
)

func RegisterCourseRouter(router gin.IRouter) {
	course := router.Group("/course")
	{
		course.POST("/", controllers.UploadCourse)
	}
}
