package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type TagsController struct{}

func (this *TagsController) Show(c *gin.Context) {
	opts := *params(c, "tag")
	service := models.NewTagService(opts)
	if tag, ok := service.FindTag(); ok {
		c.JSON(200, tag)
	} else {
		c.JSON(404, "Tag not found")
	}
}

func (this *TagsController) Index(c *gin.Context) {
	opts := *params(c)
	service := models.NewTagService(opts)
	tags := service.FindTags()
	c.JSON(200, tags)
}
