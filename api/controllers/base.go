package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// Move to middleware
func params(c *gin.Context, p ...string) *map[string]interface{} {
	params := make(map[string]interface{})
	for _, v := range c.Params {
		params[v.Key] = v.Value
	}
	c.Request.ParseForm()
	if query := c.Request.Form.Get("q"); query != "" {
		params["query"] = query
	}
	if limit := c.Request.Form.Get("limit"); limit != "" {
		if limit, err := strconv.Atoi(limit); err == nil {
			params["limit"] = limit
		}
	}
	if page := c.Request.Form.Get("page"); page != "" {
		if page, err := strconv.Atoi(page); err == nil {
			params["page"] = page
		}
	}
	for _, param := range p {
		if v := c.Request.Form.Get(param); v != "" {
			if params[param] == nil {
				params[param] = v
			}
		}
	}
	return &params
}
