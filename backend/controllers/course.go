package controllers

import (
	"github.com/gin-gonic/gin"
	"students/proto"
	"students/services"
	"students/utils"
)

func UploadCourse(c *gin.Context) {
	var req proto.UploadCourseRequest
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	resp, err := services.UploadCourseHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}