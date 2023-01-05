//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"test/controllers"
	"test/usecases"
)

func InitializeApp() *App {

	wire.Build(
		NewApp,
		NewHttpClient,
		usecases.NewPingExchangesUseCase,
		controllers.NewPingController,
	)

	return &App{}
}
