package utils

import (
	"github.com/sirupsen/logrus"
	"learning/pkg/context"
	"learning/proto"
	"net/http"
)

const (
	ErrorCode = 100
)

func Error(c *context.Context, code int, err error) {
	logrus.WithFields(logrus.Fields{
		"code": code,
		"err":  err,
		"msg":  err.Error(),
	})
	c.JSON(http.StatusOK, proto.BaseResp{
		Code:    code,
		Message: err.Error(),
	})
}

func Success(c *context.Context, data interface{}) {
	c.JSON(http.StatusOK, proto.BaseResp{
		Code: http.StatusOK,
		Data: data,
	})
}

func OK(c *context.Context) {
	c.JSON(http.StatusOK, proto.BaseResp{
		Code:    http.StatusOK,
		Message: "Success",
	})
}
