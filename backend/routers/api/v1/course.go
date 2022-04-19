package v1

import (
	"github.com/gin-gonic/gin"
	"learning/controllers"
	"learning/pkg/context"
)

func RegisterCourseRouter(router gin.IRouter) {
	course := router.Group("/course")
	{
		course.POST("/", context.WrapperHandler(controllers.UploadCourse))
	}
}
