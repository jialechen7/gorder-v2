package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jialechen7/gorder-v2/common/tracing"
)

type BaseResponse struct{}
type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	TraceID string `json:"trace_id"`
}

func (base *BaseResponse) Response(c *gin.Context, err error, data any) {
	if err != nil {
		base.Error(c, err)
	} else {
		base.Success(c, data)
	}
}

func (base *BaseResponse) Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, response{
		Code:    0,
		Message: "success",
		Data:    data,
		TraceID: tracing.TraceID(c.Request.Context()),
	})
}

func (base *BaseResponse) Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, response{
		Code:    2,
		Message: err.Error(),
		Data:    nil,
		TraceID: tracing.TraceID(c.Request.Context()),
	})
}
