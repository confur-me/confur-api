package server

import (
	//"bytes"
	"github.com/confur-me/confur-api/lib/config"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Engine *gin.Engine
}

func (this *Application) Run() {
	// TODO: use file output for logger
	//gin.DefaultWriter = bytes.NewWriter("development.log")
	this.Engine = gin.Default()
	router := Router{this.Engine}
	router.Initialize()

	this.Engine.Run(
		config.GetString("host") +
			":" +
			config.GetString("port"))
}
