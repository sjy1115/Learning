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

	isTeacher := c.UserToken.Role == consts.RoleTeacher

	exercises, err := dao.ExercisesGetByChapterId(c.Ctx, req.ChapterId)
	if err != nil {
		return nil, err
	}
	resp.ExercisesId = int64(exercises.Id)

	exerciseItems, err := dao.ExerciseItemGetByExerciseId(c.Ctx, exercises.Id)
	if err != nil {
		return nil, err
	}
	resp.Total = int64(len(exerciseItems))

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

		if isTeacher {
			item.Answer = exerciseItem.Answer
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

	exercise, err := dao.ExercisesGetByChapterId(c.Ctx, req.ChapterId)
	if err != nil {
		return nil, err
	}

	tx := mysql.GetRds(c.Ctx).Begin()

	err = dao.ExercisesUpdateById(c.Ctx, exercise.Id, map[string]interface{}{
		"title":      req.Title,
		"attachment": req.Attachment,
		"update_tm":  time.Now(),
	}, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = dao.ExercisesItemDeleteByExerciseId(c.Ctx, exercise.Id, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var exerciseItems []models.ExerciseItem
	for _, item := range req.Questions {
		options, err := jsoniter.MarshalToString(item.Options)
		if err != nil {
			tx.Rollback()
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

	exercise, err := dao.ExercisesGetById(c.Ctx, req.ExerciseId)
	if err != nil {
		return nil, err
	}

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

	err = dao.UpdateStudentScoreByStudentId(c.Ctx, c.UserToken.UserId, exercise.ChapterId, int(resp.Score))
	if err != nil {
		return nil, err
	}

	return
}
