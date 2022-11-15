//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"go-salaries-app/app"
	"go-salaries-app/controller"
	"go-salaries-app/middleware"
	"go-salaries-app/repository"
	"go-salaries-app/service"
	"net/http"
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		repository.NewSalaryRepository,
		service.NewSalaryService,
		controller.NewSalaryController,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
