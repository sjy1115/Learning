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
		user.GET("/logout", context.WrapperHandler(controllers.UserLogout))
		user.POST("/register", context.WrapperHandler(controllers.UserRegister))
		user.GET("/verifycode", context.WrapperHandler(controllers.VerifyCode))
	}
}
