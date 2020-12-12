package request

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	// 请求
	request struct {
	}
	// 请求标准
	IRequest interface {
		// 绑定
		Bind(c *gin.Context, obj interface{}) error
	}
)

const (
	errUrlParam   = "url param not exist: %v\n"
	errParseParam = "url param parse err: %w\n"
	errParamType  = "url param type error: %v\n"
)

var (
	_request *request
	once     sync.Once
)

func New() IRequest {
	once.Do(func() {
		_request = &request{}
	})
	return _request
}

func Bind(c *gin.Context, obj interface{}) error {
	New()
	return _request.Bind(c, obj)
}

func (r *request) Bind(c *gin.Context, obj interface{}) error {
	rs := reflect.TypeOf(obj).Elem()
	rt := reflect.ValueOf(obj).Elem()
	for i := 0; i < rs.NumField(); i++ {
		param := rs.Field(i).Tag.Get("param")
		if len(param) == 0 {
			continue
		}
		urlParam, ok := c.Params.Get(param)
		if !ok {
			return fmt.Errorf(errUrlParam, param)
		}

		var value interface{}
		switch rs.Field(i).Type.Kind() {
		case reflect.String:
			value = urlParam
		case reflect.Uint64:
			v, err := strconv.ParseUint(urlParam, 10, 64)
			if nil != err {
				return fmt.Errorf(errParseParam, err)
			}
			value = v
		case reflect.Uint32:
			v, err := strconv.ParseUint(urlParam, 10, 32)
			if nil != err {
				return fmt.Errorf(errParseParam, err)
			}
			value = v
		default:
			return fmt.Errorf(errParamType, param)
		}

		rt.Field(i).Set(reflect.ValueOf(value))
	}
	bType := binding.Default(c.Request.Method, c.ContentType())
	if binding.JSON == bType {
		err := c.ShouldBindBodyWith(obj, binding.JSON)
		if nil != err {
			return err
		}
	} else {
		err := c.ShouldBindWith(obj, bType)
		if nil != err {
			return err
		}
	}

	return nil
}
