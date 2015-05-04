package server

import (
	"github.com/confur-me/confur-api/lib/config"
	_ "github.com/confur-me/confur-api/lib/logrus"
	"github.com/gin-gonic/gin"
	"github.com/stephenmuss/ginerus"
)

type Application struct {
	Engine *gin.Engine
}

func (this *Application) Run() {
	this.Engine = gin.New()
	this.Engine.Use(ginerus.Ginerus())
	this.Engine.Use(gin.Recovery())

	router := Router{this.Engine}
	router.Initialize()

	this.Engine.Run(
		config.Config().UString("host", "0.0.0.0") +
			":" +
			config.Config().UString("port", "8080"))
}
