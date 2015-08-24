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
	status := 200
	opts := *params(c, "conference", "tag", "event", "shuffle")
	service := models.NewVideoService(opts)
	if videos, count, limit, offset, err := service.Videos(); err == nil {
		writeRangeHeader(c, count, limit, offset)
		if count == 0 {
			status = 204
		} else if count > limit {
			status = 206 // Partial request status
		}
		c.JSON(status, videos)
	} else {
		c.JSON(500, err)
	}
}
