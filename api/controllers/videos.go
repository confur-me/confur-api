package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type VideosController struct{}

func (this *VideosController) Show(c *gin.Context) {
	opts := *params(c, "video")
	service := models.NewVideoService(opts)
	if video, ok := service.FindVideo(); ok {
		c.JSON(200, video)
	} else {
		c.JSON(404, "Video not found")
	}
}

func (this *VideosController) Index(c *gin.Context) {
	opts := *params(c, "conference", "tag", "event")
	service := models.NewVideoService(opts)
	conferences := service.FindVideos()
	c.JSON(200, conferences)
}
