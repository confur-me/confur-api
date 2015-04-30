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

	eventsController := &controllers.EventsController{}
	this.engine.GET("/conferences/:id/events", eventsController.Index)

	videosController := &controllers.VideosController{}
	this.engine.GET("/conferences/:id/videos", videosController.Index)
	this.engine.GET("/videos/tag/:tag", videosController.ByTag)

	tagsController := &controllers.TagsController{}
	this.engine.GET("/tags", tagsController.Index)
	this.engine.GET("/tags/:slug", tagsController.Show)

	//this.engine.GET("/users/sign_in", usersController.SignIn)
	//this.engine.POST("/users/sign_in", usersController.SignIn)
	//this.engine.GET("/users/sign_up", usersController.SignIn)
	//this.engine.POST("/users/sign_up", usersController.SignIn)
	//this.engine.DELETE("/users/sign_out", usersController.SignOut)

}
