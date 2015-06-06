package server

import (
	"github.com/confur-me/confur-api/api/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func (this *Router) Initialize() {
	api := this.engine.Group("/api")
	{
		apiController := &controllers.ApiController{}
		videosController := &controllers.VideosController{}
		conferencesController := &controllers.ConferencesController{}
		eventsController := &controllers.EventsController{}
		tagsController := &controllers.TagsController{}
		searchController := &controllers.SearchController{}

		api.GET("/", apiController.Index)
		api.GET("/status", apiController.Status)

		api.GET("/conferences", conferencesController.Index)
		api.GET("/conferences/:conference", conferencesController.Show)
		api.GET("/conferences/:conference/events", eventsController.Index)
		api.GET("/conferences/:conference/videos", videosController.Index)

		api.GET("/events", eventsController.Index)
		api.GET("/events/:event", eventsController.Show)
		api.GET("/events/:event/videos", videosController.Index)

		api.GET("/videos", videosController.Index)
		api.GET("/videos/:video", videosController.Show)

		api.GET("/tags", tagsController.Index)
		api.GET("/tags/:tag", tagsController.Show)
		api.GET("/tags/:tag/videos", videosController.Index)

		api.GET("/search", searchController.Index)
	}
}
