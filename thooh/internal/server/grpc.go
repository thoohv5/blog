package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	pbarticle "thooh/api/blog/v1/article"
	pbcategory "thooh/api/blog/v1/category"
	pbcommand "thooh/api/common/v1/command"
	pbuser "thooh/api/common/v1/user"
	pbbase "thooh/api/wechat/v1/base"
	"thooh/internal/conf"
	"thooh/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	logger log.Logger,
	categoryService *service.CategoryService,
	articleService *service.ArticleService,
	commandService *service.CommandService,
	chatService *service.WeChatService,
	userService *service.UserService,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pbcategory.RegisterCategoryServer(srv, categoryService)
	pbarticle.RegisterArticleServer(srv, articleService)
	pbcommand.RegisterCommandServer(srv, commandService)
	pbbase.RegisterWeChatServer(srv, chatService)
	pbuser.RegisterUserServer(srv, userService)
	return srv
}
