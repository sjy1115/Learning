package controllers

import (
	"learning/proto"
	"learning/services"
	"learning/utils"

	"github.com/gin-gonic/gin"
)

//func AvatarUpload(c *gin.Context) {
//	var req proto.AvatarUploadReq
//	c.Bind(&req)
//
//	resp, err := services.AvatarUploadHandler(c, &req)
//	if err != nil {
//		utils.Error(c, err)
//		return
//	}
//
//	utils.Success(c, resp)
//}

func StaticDownload(c *gin.Context) {
	var req proto.AvatarDownloadReq
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	resp, err := services.StaticDownloadHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}
