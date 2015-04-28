package app

import (
	"github.com/confur-me/confur-api/app/controllers"
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
	this.engine.POST("/conferences", conferencesController.Create)
	this.engine.PUT("/conferences/:id", conferencesController.Update)
	this.engine.PATCH("/conferences/:id", conferencesController.Update)
	this.engine.DELETE("/conferences/:id", conferencesController.Destroy)
}
