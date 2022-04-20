package utils

import (
	"learning/pkg/context"
	"learning/proto"
	"net/http"
)

const (
	ErrorCode = 100
)

func Error(c *context.Context, code int, err error) {
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
