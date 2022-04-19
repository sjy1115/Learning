package utils

import (
	"learning/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, proto.BaseResp{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, proto.BaseResp{
		Code: http.StatusOK,
		Data: data,
	})
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, proto.BaseResp{
		Code:    http.StatusOK,
		Message: "Success",
	})
}
