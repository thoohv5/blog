package markdown

import "html/template"

// 请求
type Req struct {
	// label
	Label string `form:"label" json:"label" param:"label"`
}

// 返回值
type Resp struct {
	List    []string
	Name    string
	Content template.HTML
}
