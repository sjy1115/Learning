package proto

type ExercisesListRequest struct {
	ChapterId int `json:"chapter_id" form:"chapter_id"`
	Page      int `json:"page" form:"page"`
	PageSize  int `json:"page_size" form:"page_size"`
}

type ExercisesListResponse struct {
	ExercisesId int64                `json:"exercises_id"`
	Total       int64                `json:"total"`
	Items       []*ExercisesListItem `json:"items"`
}

type ExercisesListItem struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Answer  string   `json:"answer"`
	Options []string `json:"options"`
}

type ExercisesCreateRequest struct {
	ChapterId  int              `json:"chapter_id" form:"chapter_id"`
	Attachment string           `json:"attachment" form:"attachment"`
	Title      string           `json:"title" form:"title"`
	Questions  []*ExercisesItem `json:"questions" form:"questions"`
}

type ExercisesItem struct {
	Title   string   `json:"title" form:"title"`
	Type    string   `json:"type" form:"type"`
	Options []string `json:"options" form:"options"`
	Answer  string   `json:"answer" form:"answer"`
}

type ExercisesCreateResponse struct {
	Id int64 `json:"id"`
}

type ExerciseDetailRequest struct{}

type ExerciseDetailResponse struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}

type ExercisesCheckRequest struct {
	ExerciseId int            `json:"exercise_id" form:"exercise_id"`
	Answers    map[int]string `json:"answers" form:"answers"`
}

type ExercisesCheckResponse struct {
	Score int64 `json:"score"`
}
