package controllers

import (
	"students/proto"
	"students/services"
	"students/utils"

	"github.com/gin-gonic/gin"
)

func StudentList(c *gin.Context) {
	var req proto.StudentListReq
	c.Bind(&req)

	resp, err := services.StudentsListHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}

func StudentCreate(c *gin.Context) {
	var req proto.StudentCreateReq
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, err)
	}

	resp, err := services.StudentCreateHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}

func StudentDetail(c *gin.Context) {
	var req proto.StudentDetailReq
	c.Bind(&req)

	resp, err := services.StudentDetailHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}

func StudentUpdate(c *gin.Context) {
	var req proto.StudentUpdateReq
	c.Bind(&req)

	resp, err := services.StudentUpdateHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}

func StudentDelete(c *gin.Context) {
	var req proto.StudentDeleteReq
	c.Bind(&req)

	err := services.StudentDeleteHander(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.OK(c)
}
