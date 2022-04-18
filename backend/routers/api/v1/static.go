package v1

import (
	"students/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterStaticRouter(group gin.IRouter) {
	static := group.Group("/static")
	{
		static.GET("/*filename", controllers.AvatarDownload)
		static.POST("", controllers.AvatarUpload)
	}
}
