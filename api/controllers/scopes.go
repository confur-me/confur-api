package controllers

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/gin-gonic/gin"
)

type ScopesController struct{}

func (this *ScopesController) Show(c *gin.Context) {
	params := *params(c, "scope")
	service := models.NewScopeService(params)
	if scope, err := service.Scope(); err == nil {
		c.JSON(200, &scope)
	} else {
		c.JSON(404, err)
	}
}

func (this *ScopesController) Index(c *gin.Context) {
	params := *params(c)
	service := models.NewScopeService(params)
	if scopes, err := service.Scopes(); err == nil {
		c.JSON(200, &scopes)
	} else {
		c.JSON(500, err)
	}
}
