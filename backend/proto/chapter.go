package proto

type ChapterListRequest struct {
	CourseId int    `json:"id" form:"id"` // 课程id
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type ChapterListResponseItem struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	LearnNum     int64  `json:"learn_num"`
	SignNum      int64  `json:"sign_num"`
	Total        int64  `json:"total"`
	PdfUrl       string `json:"pdf_url"`
	VideoUrl     string `json:"video_url"`
	CreateAt     int64  `json:"create_at"`
	SignIn       int    `json:"sign_in"`
	PostSignIn   int    `json:"post_sign_in"`
	Learned      bool   `json:"learned"`
	Score        int64  `json:"score"`
}

type ChapterListResponse struct {
	Total int64                      `json:"total"`
	Items []*ChapterListResponseItem `json:"items"`
}

type ChapterUpdateRequest struct {
	Id           int    `json:"id" form:"id"`
	CourseId     int    `json:"course_id" form:"course_id"`
	Name         string `json:"name" form:"name"`
	Introduction string `json:"introduction" form:"introduction"`
}

type ChapterUpdateResponse struct {
	Id int `json:"id"`
}

type ChapterDetailRequest struct {
	ID int `json:"id" form:"id"`
}

type ChapterDetailResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Learned      int64  `json:"learned"`
	Total        int64  `json:"total"`
	PdfUrl       string `json:"pdf_url"`
	VideoUrl     string `json:"video_url"`
	CreateAt     int64  `json:"create_at"`
}

type ChapterCreateRequest struct {
	CourseId     int    `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Introduction string `json:"introduction" form:"introduction"`
	Pdf          string `json:"pdf" form:"pdf"`
	Video        string `json:"video" form:"video"`
}

type ChapterLearnRequest struct {
	ChapterId int `json:"chapter_id" form:"chapter_id"`
}

type ChapterLearnResponse struct{}

type ChapterCreateResponse struct {
	Id int `json:"id"`
}

type ChapterDeleteRequest struct {
	Id int `json:"id" form:"id"`
}

type ChapterDeleteResponse struct {
}

type PostSignInRequest struct {
	ChapterId int `json:"chapter_id" form:"chapter_id"`
}

type PostSignInResponse struct {
	Code string `json:"code"`
}

type SignInRequest struct {
	ChapterId int    `json:"chapter_id" form:"chapter_id"`
	Code      string `json:"code" form:"code"`
}

type SignInResponse struct {
	Status string `json:"status"`
}
