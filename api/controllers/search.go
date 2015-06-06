package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type SearchController struct{}

func (this *SearchController) Index(c *gin.Context) {
	params := *params(c)
	service := models.NewSearchService(params)
	if results, err := service.Search(); err == nil {
		c.JSON(200, &results)
	} else {
		c.JSON(500, err)
	}
}
