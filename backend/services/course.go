package services

import (
	"learning/pkg/context"
	"learning/pkg/oss"
	"learning/proto"
)

func UploadCourseHandler(c *context.Context, req *proto.UploadCourseRequest) (resp interface{}, err error) {
	//fileName := req.Filename

	fh, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	file, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = oss.Bucket.PutReader(fh.Filename, file, fh.Size)
	if err != nil {
		return nil, err
	}

	return
}
