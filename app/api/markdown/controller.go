package markdown

import (
	"github.com/gin-gonic/gin"

	"github.com/thoohv5/blog/library/request"
	"github.com/thoohv5/blog/library/response"
)

type controller struct {
	markdownDir string
}

func New() *controller {
	return &controller{
		markdownDir: "markdown",
	}
}

func (s *controller) Markdown(gtx *gin.Context) {

	req := new(Req)
	if err := request.Bind(gtx, req); nil != err {
		response.Error(gtx, err)
		return
	}

	ret, err := NewService().Markdown(&Req{Label: req.Label})
	if nil != err {
		response.Error(gtx, err)
		return
	}

	response.Success(gtx, ret, "index.tmpl")
}
