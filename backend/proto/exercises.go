package proto

type ExercisesListRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type ExercisesListResponse struct {
	Total int64                `json:"total"`
	Items []*ExercisesListItem `json:"items"`
}

type ExercisesListItem struct {
}

type ExercisesCreateRequest struct {
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
