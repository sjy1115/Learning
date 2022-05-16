package proto

type CourseListRequest struct {
	Semester string `json:"semester" form:"semester"`
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type CourseListResponseItem struct {
	ID            int    `json:"id"`
	Name          string `json:"name,omitempty"`
	Semester      string `json:"semester,omitempty"`
	Teacher       string `json:"teacher,omitempty"`
	TeacherAvatar string `json:"teacher_avatar,omitempty"`
	StudentNum    int64  `json:"student_num,omitempty"`
	CreateTm      int64  `json:"create_tm,omitempty"`
	InviteCode    string `json:"invite_code,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	Introduction  string `json:"introduction,omitempty"`
}

type CourseUpdateRequest struct {
	Name         string `json:"name" form:"name"`
	Introduction string `json:"introduction" form:"introduction"`
	Semester     string `json:"semester" form:"semester"`
}

type CourseUpdateResponse struct {
	Id int `json:"id"`
}

type CourseCreateRequest struct {
	Name         string `json:"name" form:"name"`
	Introduction string `json:"introduction" form:"introduction"`
	Semester     string `json:"semester" form:"semester"`
	Avatar       string `json:"avatar"`
}

type CourseCreateResponse struct {
	Id int `json:"id"`
}

type CourseDetailRequest struct {
	ID int `json:"id" form:"id"`
}

type CourseDetailResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Semester     string `json:"semester"`
	Teacher      string `json:"teacher"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	//StudentNum int64  `json:"student_num"`
	CreateTm int64 `json:"create_tm"`
}

type CourseListResponse struct {
	Total int64                     `json:"total"`
	Items []*CourseListResponseItem `json:"items"`
}

type CourseDeleteRequest struct {
	ID int `json:"id" form:"id"`
}

type CourseDeleteResponse struct {
}

type UploadCourseRequest struct {
	Filename string `json:"filename" form:"filename"`
}

type StartChatRequest struct {
	Token    string `json:"token" form:"token"`
	CourseId int    `json:"course_id" form:"course_id"`
}

type JoinCourseRequest struct {
	InviteCode string `json:"invite_code" form:"invite_code"`
}

type JoinCourseResponse struct{}
