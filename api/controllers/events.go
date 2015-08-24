package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type EventsController struct{}

func (this *EventsController) Show(c *gin.Context) {
	opts := *params(c, "event")
	service := models.NewEventService(opts)
	if event, err := service.Event(); err == nil {
		c.JSON(200, &event)
	} else {
		c.JSON(404, err)
	}
}

func (this *EventsController) Index(c *gin.Context) {
	opts := *params(c, "conference", "shuffle")
	service := models.NewEventService(opts)
	if events, err := service.Events(); err == nil {
		c.JSON(200, events)
	} else {
		c.JSON(500, err)
	}
}
