package v1

import (
	"learning/controllers"
	"learning/pkg/context"

	"github.com/gin-gonic/gin"
)

func RegisterStaticRouter(group gin.IRouter) {
	static := group.Group("/static")
	{
		static.GET("/*path", context.WrapperHandler(controllers.StaticDownload))
		//static.POST("", controllers.AvatarUpload)
	}
}
