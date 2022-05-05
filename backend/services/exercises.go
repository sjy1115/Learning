package services

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"learning/consts"
	"learning/dao"
	"learning/db/mysql"
	"learning/models"
	"learning/pkg/context"
	"learning/proto"
	"time"
)

func ExercisesListHandler(c *context.Context, req *proto.ExercisesListRequest) (resp *proto.ExercisesListResponse, err error) {
	resp = &proto.ExercisesListResponse{
		Items: make([]*proto.ExercisesListItem, 0),
	}

	exercises, err := dao.ExercisesGetById(c.Ctx, req.ChapterId)
	if err != nil {
		return nil, err
	}

	exerciseItems, err := dao.ExerciseItemGetByExerciseId(c.Ctx, exercises.Id)
	if err != nil {
		return nil, err
	}

	for _, exerciseItem := range exerciseItems {
		item := &proto.ExercisesListItem{
			Id:    exerciseItem.Id,
			Title: exerciseItem.Title,
			Type:  exerciseItem.Type,
		}
		var options []string
		if exerciseItem.Type == consts.ExerciseTypeMultipleChoice {
			err = jsoniter.UnmarshalFromString(exerciseItem.Options, &options)
			if err != nil {
				return nil, err
			}
			item.Options = options
		}

		resp.Items = append(resp.Items, item)
	}

	return
}

func ExercisesCreateHandler(c *context.Context, req *proto.ExercisesCreateRequest) (resp *proto.ExercisesCreateResponse, err error) {
	if c.UserToken.Role != consts.RoleTeacher {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.ExercisesCreateResponse{}

	tx := mysql.GetRds(c.Ctx).Begin()

	exercise := models.Exercises{
		Title:      req.Title,
		ChapterId:  req.ChapterId,
		Attachment: req.Attachment,
		InsertTm:   time.Now(),
		UpdateTm:   time.Now(),
	}

	err = dao.Create(c.Ctx, &exercise, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var exerciseItems []models.ExerciseItem
	for _, item := range req.Questions {
		options, err := jsoniter.MarshalToString(item.Options)
		if err != nil {
			return nil, err
		}

		exerciseItem := models.ExerciseItem{
			ExerciseId: exercise.Id,
			Title:      item.Title,
			Type:       item.Type,
			Options:    options,
			Answer:     item.Answer,
			InsertTm:   time.Now(),
			UpdateTm:   time.Now(),
		}
		exerciseItems = append(exerciseItems, exerciseItem)
	}
	err = dao.Create(c.Ctx, &exerciseItems, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	resp.Id = int64(exercise.Id)

	return
}

func ExercisesCheckHandler(c *context.Context, req *proto.ExercisesCheckRequest) (resp *proto.ExercisesCheckResponse, err error) {
	resp = &proto.ExercisesCheckResponse{}

	exerciseItems, err := dao.ExerciseItemGetByExerciseId(c.Ctx, req.ExerciseId)
	if err != nil {
		return nil, err
	}

	resp.Score = 100

	for _, exerciseItem := range exerciseItems {
		if exerciseItem.Answer != req.Answers[exerciseItem.Id] {
			resp.Score -= 2
		}
	}

	if resp.Score < 0 {
		resp.Score = 0
	}

	return
}
