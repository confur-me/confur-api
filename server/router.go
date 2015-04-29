package server

import (
	"github.com/confur-me/confur-api/api/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func (this *Router) Initialize() {
	apiController := &controllers.ApiController{}
	this.engine.GET("/api", apiController.Index)

	conferencesController := &controllers.ConferencesController{}
	this.engine.GET("/conferences", conferencesController.Index)
	this.engine.GET("/conferences/:id", conferencesController.Show)
}
