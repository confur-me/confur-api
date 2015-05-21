package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type ConferencesController struct{}

func (this *ConferencesController) Show(c *gin.Context) {
	params := *params(c, "conference")
	service := models.NewConferenceService(params)
	if conference, err := service.Conference(); err == nil {
		c.JSON(200, &conference)
	} else {
		c.JSON(404, err)
	}
}

func (this *ConferencesController) Index(c *gin.Context) {
	params := *params(c)
	service := models.NewConferenceService(params)
	if conferences, err := service.Conferences(); err == nil {
		c.JSON(200, &conferences)
	} else {
		c.JSON(400, err)
	}
}
