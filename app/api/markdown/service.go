package markdown

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/thoohv5/blog/app/service/file"
	"github.com/thoohv5/blog/app/service/markdown"
)

type (
	service struct {
		markdownDir string
	}
	IService interface {
		Markdown(req *Req) (*Resp, error)
	}
)

func NewService() IService {
	return &service{
		markdownDir: "markdown",
	}
}

func (s *service) Markdown(req *Req) (*Resp, error) {
	resp := new(Resp)
	name := req.Label

	fileName := name
	if params := strings.Split(name, "."); len(params) > 0 {
		fileName = params[0]
	}
	if fileNames := strings.Split(fileName, "-"); len(fileNames) > 0 {
		fileName = strings.Join(fileNames, "/")
	}

	// content
	content, err := markdown.New().Parse(fmt.Sprintf("./%s/%s.md", s.markdownDir, fileName))
	if nil != err {
		return resp, err
	}

	// list
	files, err := file.New().Walk(fmt.Sprintf("./%s", s.markdownDir))
	if nil != err {
		return resp, err
	}
	list := make([]string, len(files))
	for _, f := range files {
		if !strings.HasSuffix(f.Name, ".md") {
			continue
		}
		list = append(list, fmt.Sprintf("%s.html", strings.ReplaceAll(strings.Trim(strings.Split(f.Name, s.markdownDir)[1], "/.md"), "/", "-")))
	}
	resp.Name = fileName
	resp.Content = template.HTML(content)
	resp.List = list
	return resp, nil
}
