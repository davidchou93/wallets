package router

import (
	"github.com/davidchou93/wallets/config"
	"github.com/davidchou93/wallets/internal/handler"
	"github.com/davidchou93/wallets/internal/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var WireSet = wire.NewSet(NewRouter)

func NewRouter(appEnv config.AppEnv, logger *logrus.Logger, adminHandler *handler.AdminHandler) *gin.Engine {
	if appEnv == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logging.GinLogger(logger))
	router.GET("", adminHandler.Echo)
	return router
}
