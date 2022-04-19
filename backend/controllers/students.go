package controllers

import (
	"learning/pkg/context"
	"learning/proto"
	"learning/services"
	"learning/utils"
)

func StudentList(c *context.Context) {
	var req proto.StudentListReq
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	resp, err := services.StudentsListHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}

func StudentCreate(c *context.Context) {
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

func StudentDetail(c *context.Context) {
	var req proto.StudentDetailReq
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	resp, err := services.StudentDetailHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}

func StudentUpdate(c *context.Context) {
	var req proto.StudentUpdateReq
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	resp, err := services.StudentUpdateHandler(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Success(c, resp)
}

func StudentDelete(c *context.Context) {
	var req proto.StudentDeleteReq
	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	err = services.StudentDeleteHander(c, &req)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.OK(c)
}
