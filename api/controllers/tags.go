package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type TagsController struct{}

// @Title Show
// @Description get Tag by slug
// @Param	slug		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tag
// @Failure 403 :slug is empty
// @router /:slug [get]
func (this *TagsController) Show(c *gin.Context) {
	slug := c.Params.ByName("slug")
	tag := models.TagBySlug(slug)
	if tag.ID > 0 {
		c.JSON(200, tag)
	} else {
		c.JSON(404, "Tag not found")
	}
}

// @Title Index
// @Description get Tag
// @Param	slug		path 	string	true
// @Success 200 {object} models.Tag
// @Failure 403
// @router / [get]
func (this *TagsController) Index(c *gin.Context) {
	tags := models.Tags()
	c.JSON(200, tags)
}
