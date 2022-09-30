//go:build wireinject
// +build wireinject

package main

import (
	"github.com/davidchou93/wallets/config"
	"github.com/davidchou93/wallets/internal/handler"
	"github.com/davidchou93/wallets/internal/logging"
	"github.com/davidchou93/wallets/internal/model"
	"github.com/davidchou93/wallets/internal/repository"
	"github.com/davidchou93/wallets/internal/repository/db"
	"github.com/davidchou93/wallets/internal/router"
	"github.com/davidchou93/wallets/internal/service"
	"github.com/google/wire"
)

func InitializeApp() (App, error) {
	wire.Build(
		logging.WireSet,
		db.WireSet,
		config.WireSet,
		repository.WireSet,
		router.WireSet,
		handler.WireSet,
		service.WireSet,
		wire.Bind(new(model.Repository), new(*repository.RepositoryImpl)),
		wire.Bind(new(model.UserService), new(*service.UserServiceImpl)),
		wire.Struct(new(App), "*"),
	)
	return App{}, nil

}
