package common

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jiahuipaung/Codefolio_Backend/common/handler/errors"
	"github.com/jiahuipaung/Codefolio_Backend/common/tracing"
	"net/http"
)

type BaseResponse struct {
}

type response struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	TraceID string `json:"trace_id"`
}

func (base *BaseResponse) success(c *gin.Context, data interface{}) {
	errno, errMsg := errors.Output(err)
	r := response{
		Errno:   0,
		Message: "success",
		Data:    data,
		TraceID: tracing.TraceID(c.Request.Context()),
	}
	c.JSON(http.StatusOK, r)
	resp, _ := json.Marshal(r)
	c.Set("response", resp)
}

func (base *BaseResponse) error(c *gin.Context, err error) {
	errno, errmsg := errors.Output(err)
	r := response{
		Errno:   errno,
		Message: errmsg,
		Data:    nil,
		TraceID: tracing.TraceID(c.Request.Context()),
	}
	c.JSON(http.StatusOK, r)
	resp, _ := json.Marshal(r)
	c.Set("response", resp)
}

func (base *BaseResponse) Response(c *gin.Context, err error, data interface{}) {
	if err != nil {
		base.error(c, err)
	} else {
		base.success(c, data)
	}
}
