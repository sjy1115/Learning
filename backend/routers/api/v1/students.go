package v1

import (
	"learning/controllers"
	"learning/pkg/context"

	"github.com/gin-gonic/gin"
)

func RegisterStudentsRouter(group gin.IRouter) {
	students := group.Group("/students")
	{
		students.GET("/", context.WrapperHandler(controllers.StudentList))
		students.POST("/", context.WrapperHandler(controllers.StudentCreate))
		students.GET("/:id", context.WrapperHandler(controllers.StudentDetail))
		students.PUT("/:id", context.WrapperHandler(controllers.StudentUpdate))
		students.DELETE("/:id", context.WrapperHandler(controllers.StudentDelete))
	}
}
