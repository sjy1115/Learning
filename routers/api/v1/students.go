package v1

import (
	"students/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterStudentsRouter(group gin.IRouter) {
	students := group.Group("/students")
	{
		students.GET("/", controllers.StudentList)
		students.POST("/", controllers.StudentCreate)
		students.GET("/:id", controllers.StudentDetail)
		students.PUT("/:id", controllers.StudentUpdate)
		students.DELETE("/:id", controllers.StudentDelete)
	}
}
