package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type EventsController struct{}

// @Title Show
// @Description get Event by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Event
// @Failure 403 :id is empty
// @router /:id [get]
func (this *EventsController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	event := models.EventById(id)
	if event.ID > 0 {
		c.JSON(200, event)
	} else {
		c.JSON(404, "Event not found")
	}
}

// @Title Index
// @Description get Event
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Event
// @Failure 403
// @router / [get]
func (this *EventsController) Index(c *gin.Context) {
	conference_id := c.Params.ByName("id")
	events := models.EventsByConference(conference_id)
	if len(events) == 0 {
		events = make([]models.Event, 0)
	}
	c.JSON(200, events)
}
