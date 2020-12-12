package response

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/thoohv5/blog/app/api/common"
)

type (
	response struct {
		ctx *gin.Context
	}
	IResponse interface {
		// Success
		Success(data interface{}, msg string)
		// DefaultSuccess
		DefaultSuccess(data interface{})
		// Error
		Error(err error)
		// Generate
		Generate(data interface{}, err error)
	}
)

var (
	_response *response
	once      sync.Once
)

func New(ctx *gin.Context) IResponse {
	once.Do(func() {
		_response = &response{}
	})
	_response.ctx = ctx
	return _response
}

// 响应成功请求
func Success(ctx *gin.Context, data interface{}, msg string) {
	New(ctx).Success(data, msg)
	return
}

// 响应成功请求
func (r *response) Success(data interface{}, msg string) {
	if len(msg) == 0 {
		msg = "success"
	}

	contentType := r.ctx.GetHeader("content-type")
	switch contentType {
	case "application/json":
		// 响应数据
		r.ctx.JSON(http.StatusOK, common.ResponseEntity{
			ErrorCode: 0,
			Message:   msg,
			Data:      data,
		})
	case "html/text":
		r.ctx.HTML(http.StatusOK, msg, data)
	default:
		r.ctx.HTML(http.StatusOK, msg, data)
	}

	return
}

func DefaultSuccess(ctx *gin.Context, data interface{}) {
	New(ctx).DefaultSuccess(data)
	return
}

func (r *response) DefaultSuccess(data interface{}) {
	r.Success(data, "成功")
}

// 响应错误
func Error(ctx *gin.Context, err error) {
	New(ctx).Error(err)
	return
}

// 响应错误
func (r *response) Error(err error) {
	if err == nil {
		return
	}

	contentType := r.ctx.GetHeader("content-type")
	switch contentType {
	case "application/json":
		r.ctx.JSON(http.StatusOK, common.ResponseEntity{
			ErrorCode: 200,
			Message:   err.Error(),
			Data:      nil,
		})
	case "html/text":
		r.ctx.HTML(http.StatusNotFound, "404.tmpl", nil)
	default:
		r.ctx.HTML(http.StatusNotFound, "404.tmpl", nil)
	}

	return
}

// 通用响应
func Generate(ctx *gin.Context, data interface{}, err error) {
	New(ctx).Generate(data, err)
	return
}

// 通用响应
func (r *response) Generate(data interface{}, err error) {
	if nil != err {
		r.Error(err)
		return
	}
	r.Success(data, "success")
}
