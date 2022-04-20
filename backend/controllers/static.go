package controllers

import (
	"learning/pkg/context"
	"learning/proto"
	"learning/services"
	"learning/utils"
)

func StaticDownload(c *context.Context) {
	var req proto.DownloadReq
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
