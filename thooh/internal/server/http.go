package server

import (
	ghttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	pbarticle "thooh/api/blog/v1/article"
	pbcategory "thooh/api/blog/v1/category"
	pbcommand "thooh/api/common/v1/command"
	"thooh/internal/conf"
	"thooh/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	c *conf.Server,
	logger log.Logger,
	categoryService *service.CategoryService,
	articleService *service.ArticleService,
	commandService *service.CommandService,
) *http.Server {
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

	srv.Route("/")
	srv.HandleFunc("/wechat", serveWechat)

	pbcategory.RegisterCategoryHTTPServer(srv, categoryService)
	pbarticle.RegisterArticleHTTPServer(srv, articleService)
	pbcommand.RegisterCommandHTTPServer(srv, commandService)
	return srv
}

func serveWechat(rw ghttp.ResponseWriter, req *ghttp.Request) {
	m := make(map[string]string)
	mp := strings.FieldsFunc(req.URL.RawQuery, func(r rune) bool {
		if r == '=' || r == '&' {
			return true
		}
		return false
	})
	for i := 0; i <len(mp); i+=2  {
		m[mp[i]] = mp[i+1]
	}
	// rw.WriteHeader(200)
	rw.Write([]byte(m["echostr"]))
	return
	// wc := wechat.NewWechat()
	// // 这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	// memory := cache.NewMemory()
	// cfg := &offConfig.Config{
	// 	AppID:     "xxx",
	// 	AppSecret: "xxx",
	// 	Token:     "xxx",
	// 	// EncodingAESKey: "xxxx",
	// 	Cache: memory,
	// }
	// officialAccount := wc.GetOfficialAccount(cfg)
	//
	// // 传入request和responseWriter
	// server := officialAccount.GetServer(req, rw)
	// // 设置接收消息的处理方法
	// server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
	// 	// TODO
	// 	// 回复消息：演示回复用户发送的消息
	// 	text := message.NewText(msg.Content)
	// 	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	// })
	//
	// // 处理消息接收以及回复
	// err := server.Serve()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // 发送回复的消息
	// server.Send()
}