package v1

import (
	"github.com/gin-gonic/gin"
	"learning/controllers"
	"learning/pkg/context"
)

func RegisterStudentsRouter(group gin.IRouter) {
	user := group.Group("/user")
	{
		user.POST("/login", context.WrapperHandler(controllers.UserLogin))
		user.GET("/info", context.WrapperHandler(controllers.UserInfo))
		user.PUT("/:id", context.WrapperHandler(controllers.UserUpdate))
		user.GET("/logout", context.WrapperHandler(controllers.UserLogout))
		user.POST("/register", context.WrapperHandler(controllers.UserRegister))
		user.POST("/changepassword", context.WrapperHandler(controllers.ChangePassword))
		user.GET("/verifycode", context.WrapperHandler(controllers.VerifyCode))
	}
}
