package controllers

import (
	"learning/pkg/context"
	"learning/proto"
	"learning/services"
	"learning/utils"
)

func StaticDownload(c *context.Context) {
	var req proto.DownloadRequest
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	resp, err := services.StaticDownloadHandler(c, &req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}

func StaticUpload(c *context.Context) {
	resp, err := services.StaticUploadHandler(c)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}
