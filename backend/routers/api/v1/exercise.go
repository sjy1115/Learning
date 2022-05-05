package v1

import (
	"github.com/gin-gonic/gin"
	"learning/controllers"
	"learning/pkg/context"
)

func RegisterExerciseRouter(c gin.IRouter) {
	exercise := c.Group("/exercise")
	{
		exercise.GET("", context.WrapperHandler(controllers.ExerciseList))
		exercise.POST("", context.WrapperHandler(controllers.ExerciseCreate))
		exercise.POST("/check", context.WrapperHandler(controllers.ExerciseCheck))
	}
}
