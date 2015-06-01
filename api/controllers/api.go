package controllers

import (
	"github.com/gin-gonic/gin"
)

const VERSION string = "0.0.1"

type ApiController struct{}

func (this *ApiController) Index(c *gin.Context) {
	api := make(map[string]string)
	api["version"] = VERSION

	c.JSON(200, api)
}

func (this *ApiController) Status(c *gin.Context) {
	c.String(200, "ok")
}
