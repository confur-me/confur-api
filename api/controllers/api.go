package controllers

import (
	"github.com/confur-me/confur-api/lib/config"
	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (this *ApiController) Index(c *gin.Context) {
	api := make(map[string]string)
	api["version"] = config.Config().UString("api.version")

	c.JSON(200, api)
}
