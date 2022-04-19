package proto

type UploadCourseRequest struct {
	Filename string `json:"filename" form:"filename"`
}
