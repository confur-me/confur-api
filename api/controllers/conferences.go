package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type ConferencesController struct{}

// @Title Show
// @Description get Conference by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Conference
// @Failure 403 :id is empty
// @router /:id [get]
func (this *ConferencesController) Show(c *gin.Context) {
	slug := c.Params.ByName("slug")
	conference := models.ConferenceBySlug(slug)
	if conference.Slug != "" {
		c.JSON(200, conference)
	} else {
		c.JSON(404, "Conference not found")
	}
}

// @Title Index
// @Description get Conference
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Conference
// @Failure 403
// @router / [get]
func (this *ConferencesController) Index(c *gin.Context) {
	conferences := models.Conferences()
	c.JSON(200, conferences)
}
