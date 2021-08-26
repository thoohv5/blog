// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"thooh/internal/biz"
	"thooh/internal/conf"
	"thooh/internal/data"
	"thooh/internal/server"
	"thooh/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, assets *conf.Assets, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	categoryRepo := data.NewCategoryRepo(dataData, logger)
	categoryUsecase := biz.NewCategoryUsecase(categoryRepo, logger)
	categoryService := service.NewCategoryService(categoryUsecase)
	articleRepo := data.NewArticleRepo(dataData, logger)
	readRepo := data.NewReadRepo(dataData, logger)
	articleUsecase := biz.NewArticleUsecase(articleRepo, readRepo, logger)
	articleService := service.NewArticleService(articleUsecase)
	file := biz.NewFile()
	commandService := service.NewCommandService(file, assets, articleUsecase)
	httpServer := server.NewHTTPServer(confServer, logger, categoryService, articleService, commandService)
	grpcServer := server.NewGRPCServer(confServer, logger, categoryService, articleService, commandService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}