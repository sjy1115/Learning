package main

import (
	"context"
	"fmt"
	"learning/config"
	"learning/consts"
	"learning/db/mysql"
	"learning/models"
)

func main() {
	ctx := context.TODO()

	err := config.Init("")
	if err != nil {
		panic(err)
	}
	config.Conf.Web.Docker = false

	err = mysql.InitMysql(config.Conf)
	if err != nil {
		panic(err)
	}

	db := mysql.GetRds(ctx)

	items := []models.ExerciseItem{
		{
			Title: "1",
			Type:  consts.ExerciseTypeFillsUp,
		},
		{
			Title: "2",
			Type:  consts.ExerciseTypeMultipleChoice,
		},
		{
			Title: "3",
			Type:  consts.ExerciseTypeJudge,
		},
	}

	err = db.Create(&items).Error
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Println(item.Id)
	}
}
