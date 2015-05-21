package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type VideosController struct{}

func (this *VideosController) Show(c *gin.Context) {
	opts := *params(c, "video")
	service := models.NewVideoService(opts)
	if video, err := service.Video(); err == nil {
		c.JSON(200, video)
	} else {
		c.JSON(404, err)
	}
}

func (this *VideosController) Index(c *gin.Context) {
	opts := *params(c, "conference", "tag", "event")
	service := models.NewVideoService(opts)
	if videos, err := service.Videos(); err == nil {
		c.JSON(200, videos)
	} else {
		c.JSON(400, err)
	}
}
