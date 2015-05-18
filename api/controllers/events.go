package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type EventsController struct{}

func (this *EventsController) Show(c *gin.Context) {
	opts := *params(c, "event")
	service := models.NewEventService(opts)
	if event, ok := service.FindEvent(); ok {
		c.JSON(200, event)
	} else {
		c.JSON(404, "Event not found")
	}
}

func (this *EventsController) Index(c *gin.Context) {
	opts := *params(c, "conference")
	service := models.NewEventService(opts)
	events := service.FindEvents()
	c.JSON(200, events)
}
