package v1

import (
	"github.com/gin-gonic/gin"
	"learning/controllers"
	"learning/pkg/context"
)

func RegisterCourseRouter(router gin.IRouter) {
	course := router.Group("/course")
	{
		course.GET("", context.WrapperHandler(controllers.CourseList))
		course.GET("/:id", context.WrapperHandler(controllers.CourseDetail))
		course.POST("", context.WrapperHandler(controllers.CourseCreate))
		course.PUT("/:id", context.WrapperHandler(controllers.CourseUpdate))
		course.DELETE("/:id", context.WrapperHandler(controllers.CourseDelete))
		course.POST("/join", context.WrapperHandler(controllers.JoinCourse))
		course.POST("/:course_id/sign_in", context.WrapperHandler(controllers.StudentSignIn))
		course.POST("/upload", context.WrapperHandler(controllers.UploadCourse))
		course.GET("/chat", context.WrapperHandler(controllers.StartChat))
	}
	chapter := router.Group("/chapter")
	{
		chapter.GET("", context.WrapperHandler(controllers.ChapterList))
		chapter.GET("/:id", context.WrapperHandler(controllers.ChapterDetail))
		chapter.POST("", context.WrapperHandler(controllers.ChapterCreate))
		chapter.GET("/students", context.WrapperHandler(controllers.ListChapterStudent))
		chapter.POST("/learn", context.WrapperHandler(controllers.ChapterLearn))
		chapter.PUT("/:id", context.WrapperHandler(controllers.ChapterUpdate))
		chapter.DELETE("/:id", context.WrapperHandler(controllers.ChapterDelete))
	}
}
