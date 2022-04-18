package controllers

import (
	"students/proto"
	"students/services"
	"students/utils"

	"github.com/gin-gonic/gin"
)

func AvatarUpload(c *gin.Context) {
	var req proto.AvatarUploadReq
	c.Bind(&req)

	resp, err := services.AvatarUploadHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}

func AvatarDownload(c *gin.Context) {

}
