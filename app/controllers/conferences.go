package controllers

import (
	"github.com/gin-gonic/gin"
)

type ConferencesController struct{}

// @Title Create
// @Description create Conference
// @Param	body		body 	models.Conference	true		"body for Conference content"
// @Success 200 {int} models.Conference.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ConferencesController) Create(c *gin.Context) {

}

// @Title Show
// @Description get Conference by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Conference
// @Failure 403 :id is empty
// @router /:id [get]
func (this *ConferencesController) Show(c *gin.Context) {

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
	// TODO: Orm fetch

	//c.JSON(200, api)
}

// @Title Update
// @Description update the Conference
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Conference	true		"body for Conference content"
// @Success 200 {object} models.Conference
// @Failure 403 :id is not int
// @router /:id [put]
func (this *ConferencesController) Update(c *gin.Context) {

}

// @Title Destroy
// @Description delete the Conference
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *ConferencesController) Destroy(c *gin.Context) {

}
