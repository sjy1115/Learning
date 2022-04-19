package proto

type UploadCourseRequest struct {
	Filename string `json:"filename" form:"filename"`
}

type StartChatRequest struct {
	Token    string `json:"token" form:"token"`
	CourseId string `json:"courseId" form:"courseId"`
}
