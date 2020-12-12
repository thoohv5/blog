package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/thoohv5/blog/app/api/common"
	"github.com/thoohv5/blog/app/api/markdown"
)

func RegisterRoute(r *gin.Engine) {
	// 公共模块
	r.LoadHTMLGlob("./templates/*")
	r.StaticFS("/static", http.Dir("./static"))
	r.NoRoute(common.New().NoRoute)
	r.NoMethod(common.New().NoRoute)

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 博客
	blog := r.Group("/blog")
	{
		blog.GET(":label", markdown.New().Markdown)
	}
}
