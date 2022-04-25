package v1

import (
	"github.com/gin-gonic/gin"
	"learning/controllers"
	"learning/pkg/context"
)

func RegisterCourseRouter(router gin.IRouter) {
	course := router.Group("/course")
	{
		course.GET("/", context.WrapperHandler(controllers.CourseList))
		course.POST("/", context.WrapperHandler(controllers.CourseCreate))
		course.POST("/", context.WrapperHandler(controllers.UploadCourse))
		course.GET("/chat", context.WrapperHandler(controllers.StartChat))
	}
}
