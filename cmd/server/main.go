package main

import (
	"fmt"
	conf "github.com/davidchou93/wallets/config"
	"github.com/davidchou93/wallets/internal/model"
	"github.com/sirupsen/logrus"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type App struct {
	Config *conf.Config
	Router *gin.Engine
	Logger *logrus.Logger
	model.UserService
}

func main() {
	myApp, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	server := endless.NewServer(fmt.Sprintf("localhost:%s", myApp.Config.Port), myApp.Router)
	server.BeforeBegin = func(add string) {
		myApp.Logger.Infof("Actual pid is %d", syscall.Getpid())
		myApp.Logger.Infof("Start server on http://localhost:%s", myApp.Config.Port)
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
