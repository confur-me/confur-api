package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type VideosController struct{}

// @Title Show
// @Description get Video by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Video
// @Failure 403 :id is empty
// @router /:id [get]
func (this *VideosController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	video := models.VideoById(id)
	if video.ID > 0 {
		c.JSON(200, video)
	} else {
		c.JSON(404, "Video not found")
	}
}

// @Title Index
// @Description get Video
// @Param	id		path 	string	true
// @Success 200 {object} models.Video
// @Failure 403
// @router / [get]
func (this *VideosController) Index(c *gin.Context) {
	conference_id := c.Params.ByName("id")
	videos := models.ConferenceVideosCollection(conference_id)
	c.JSON(200, videos)
}
