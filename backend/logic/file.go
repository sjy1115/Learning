package logic

import (
	"fmt"
	"learning/consts"
	"learning/pkg/context"
	"learning/pkg/oss"
)

func UploadCourseFile(c *context.Context, courseId int, fileType string) (string, error) {
	fh, err := c.FormFile(fileType)
	if err != nil {
		return "", err
	}

	var filePath string
	switch fileType {
	case consts.FileTypeAvatar:
		filePath = fmt.Sprintf("%s/%d/%s", consts.FileTypeAvatar, courseId, fh.Filename)
	case consts.FileTypePdf:
		filePath = fmt.Sprintf("%s/%d/%s", consts.FileTypePdf, courseId, fh.Filename)
	case consts.FileTypeVideo:
		filePath = fmt.Sprintf("%s/%d/%s", consts.FileTypeVideo, courseId, fh.Filename)
	}

	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = oss.Bucket.PutReader(filePath, file, fh.Size)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
