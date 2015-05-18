package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type ConferencesController struct{}

func (this *ConferencesController) Show(c *gin.Context) {
	opts := *params(c, "conference")
	service := models.NewConferenceService(opts)
	if conference, ok := service.FindConference(); ok {
		c.JSON(200, conference)
	} else {
		c.JSON(404, "Conference not found")
	}
}

func (this *ConferencesController) Index(c *gin.Context) {
	opts := *params(c)
	service := models.NewConferenceService(opts)
	conferences := service.FindConferences()
	c.JSON(200, conferences)
}
