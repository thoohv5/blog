// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"thooh/internal/biz"
	"thooh/internal/conf"
	"thooh/internal/data"
	"thooh/internal/server"
	"thooh/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Assets, *conf.Wechat, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
	// panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
