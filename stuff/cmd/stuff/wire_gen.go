// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"stuff/api/other/user/v1"
	"stuff/internal/biz"
	"stuff/internal/conf"
	"stuff/internal/data"
	"stuff/internal/server"
	"stuff/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, userClient v1.UserClient) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger, userClient)
	if err != nil {
		return nil, nil, err
	}
	stuffRepo := data.NewStuffRepo(dataData, logger)
	stuffUsecase := biz.NewStuffUsecase(stuffRepo, logger, userClient)
	stuffService := service.NewStuffService(stuffUsecase)
	grpcServer := server.NewGRPCServer(confServer, stuffService, logger)
	categoryRepo := data.NewCategoryRepo(dataData, logger)
	categoryUsecase := biz.NewCategoryUsecase(categoryRepo, logger)
	categoryService := service.NewCategoryService(categoryUsecase)
	httpServer := server.NewHTTPServer(confServer, stuffService, categoryService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
