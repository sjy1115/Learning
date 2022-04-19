package v1

import (
	"learning/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterStaticRouter(group gin.IRouter) {
	static := group.Group("/static")
	{
		static.GET("/*path", controllers.StaticDownload)
		//static.POST("", controllers.AvatarUpload)
	}
}
