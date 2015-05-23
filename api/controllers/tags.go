package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type TagsController struct{}

func (this *TagsController) Show(c *gin.Context) {
	opts := *params(c, "tag")
	service := models.NewTagService(opts)
	if tag, err := service.Tag(); err == nil {
		c.JSON(200, &tag)
	} else {
		c.JSON(404, err)
	}
}

func (this *TagsController) Index(c *gin.Context) {
	opts := *params(c)
	service := models.NewTagService(opts)
	if tags, err := service.Tags(); err == nil {
		c.JSON(200, &tags)
	} else {
		c.JSON(500, err)
	}
}
