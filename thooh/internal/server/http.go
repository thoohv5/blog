package server

import (
	ghttp "net/http"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	pbarticle "thooh/api/blog/v1/article"
	pbcategory "thooh/api/blog/v1/category"
	pbcommand "thooh/api/common/v1/command"
	pbbase "thooh/api/wechat/v1/base"
	"thooh/internal/conf"
	"thooh/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	c *conf.Server,
	logger klog.Logger,
	categoryService *service.CategoryService,
	articleService *service.ArticleService,
	commandService *service.CommandService,
	chatService *service.WeChatService,
) *http.Server {

	log := klog.NewHelper(logger)

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	)))
	srv := http.NewServer(opts...)

	// 微信
	srv.Route("/")
	srv.HandleFunc("/we-chat", func(rw ghttp.ResponseWriter, req *ghttp.Request) {
		if err := chatService.WeChat(rw, req); nil != err {
			log.Errorf("http WeChat err,err:%v", err)
		}
	})

	pbcategory.RegisterCategoryHTTPServer(srv, categoryService)
	pbarticle.RegisterArticleHTTPServer(srv, articleService)
	pbcommand.RegisterCommandHTTPServer(srv, commandService)
	pbbase.RegisterWeChatHTTPServer(srv, chatService)
	return srv
}
