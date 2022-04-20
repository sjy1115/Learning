package controllers

import (
	"learning/pkg/context"
	"learning/proto"
	"learning/services"
	"learning/utils"
)

func UploadCourse(c *context.Context) {
	var req proto.UploadCourseRequest
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	resp, err := services.UploadCourseHandler(c, &req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}

func StartChat(c *context.Context) {
	var req proto.StartChatRequest
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	err = services.StartChatHandler(c, &req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	return
}
