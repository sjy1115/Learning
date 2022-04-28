package proto

type ChapterListRequest struct {
	CourseId int    `json:"course_id" form:"course_id"`
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type ChapterListResponseItem struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Learned      int64  `json:"learned"`
	Total        int64  `json:"total"`
	PdfUrl       string `json:"pdf_url"`
	CreateAt     int64  `json:"create_at"`
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
	CreateAt     int64  `json:"create_at"`
}

type ChapterCreateRequest struct {
	CourseId     int    `json:"course_id" form:"course_id"`
	Name         string `json:"name" form:"name"`
	Introduction string `json:"introduction" form:"introduction"`
	Pdf          string `json:"pdf" form:"pdf"`
	Video        string `json:"video" form:"video"`
}

type ChapterCreateResponse struct {
	Id int `json:"id"`
}

type ChapterDeleteRequest struct {
	Id int `json:"id" form:"id"`
}

type ChapterDeleteResponse struct {
}
