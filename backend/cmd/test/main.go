package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//ctx := context.TODO()
	//
	//err := config.Init("")
	//if err != nil {
	//	panic(err)
	//}
	//config.Conf.Web.Docker = false
	//
	//err = mysql.InitMysql(config.Conf)
	//if err != nil {
	//	panic(err)
	//}
	//
	//db := mysql.GetRds(ctx)
	//
	//items := []models.ExerciseItem{
	//	{
	//		Title: "1",
	//		Type:  consts.ExerciseTypeFillsUp,
	//	},
	//	{
	//		Title: "2",
	//		Type:  consts.ExerciseTypeMultipleChoice,
	//	},
	//	{
	//		Title: "3",
	//		Type:  consts.ExerciseTypeJudge,
	//	},
	//}
	//
	//err = db.Create(&items).Error
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, item := range items {
	//	fmt.Println(item.Id)
	//}

	test()
}

func test() {
	t := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}

	str, _ := json.Marshal(t)
	fmt.Println(string(str))
}
