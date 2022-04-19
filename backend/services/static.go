package services

import (
	"net/url"
	"path/filepath"
	"strings"
	"students/pkg/oss"
	"students/proto"

	"github.com/gin-gonic/gin"
)

func StaticDownloadHandler(c *gin.Context, req *proto.AvatarDownloadReq) (*proto.AvatarDownloadResp, error) {
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
