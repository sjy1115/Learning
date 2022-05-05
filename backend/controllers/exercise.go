package controllers

import (
	"learning/pkg/context"
	"learning/proto"
	"learning/services"
	"learning/utils"
)

func ExerciseList(c *context.Context) {
	var req proto.ExercisesListRequest

	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	resp, err := services.ExercisesListHandler(c, &req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}

func ExerciseCreate(c *context.Context) {
	var req proto.ExercisesCreateRequest

	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	resp, err := services.ExercisesCreateHandler(c, &req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}

func ExerciseCheck(c *context.Context) {
	var req proto.ExercisesCheckRequest

	err := c.Bind(&req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	resp, err := services.ExercisesCheckHandler(c, &req)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return
	}

	utils.Success(c, resp)
}
