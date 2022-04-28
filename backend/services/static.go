package services

import (
	"fmt"
	"learning/pkg/context"
	"learning/pkg/oss"
	"learning/proto"
	"learning/utils"
	"net/url"
	"path/filepath"
	"strings"
)

func StaticDownloadHandler(c *context.Context, req *proto.DownloadRequest) (*proto.AvatarDownloadResp, error) {
	path := c.Param("path")

	data, err := oss.Bucket.Content(path)
	if err != nil {
		return nil, err
	}

	if strings.HasSuffix(path, ".pdf") {
		c.Writer.Write(data)
		return nil, err
	} else if strings.HasSuffix(path, ".mp4") {
		c.Writer.Header().Set("Content-Type", "video/mp4")
		c.Writer.Write(data)
		return nil, err
	}

	c.Header("Content-Disposition", "attachment;filename="+url.QueryEscape(filepath.Base(path)))
	_, err = c.Writer.Write(data)
	if err != nil {
		return nil, err
	}
	//c.Writer.Header().Set("Content-Type", "video/mp4")
	//http.ServeContent()

	return nil, nil
}

func StaticUploadHandler(c *context.Context) (resp *proto.StaticUploadResponse, err error) {
	fh, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	resp = &proto.StaticUploadResponse{}

	file, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	randomStr := utils.RandomString(32)
	filePath := fmt.Sprintf("%s-%s", randomStr, fh.Filename)

	err = oss.Bucket.PutReader(filePath, file, fh.Size)
	if err != nil {
		return nil, err
	}

	resp.Path = filePath

	return
}
