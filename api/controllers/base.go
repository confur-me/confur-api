package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func params(c *gin.Context, params ...string) *map[string]interface{} {
	opts := make(map[string]interface{})
	for _, v := range c.Params {
		opts[v.Key] = v.Value
	}
	c.Request.ParseForm()
	if query := c.Request.Form.Get("q"); query != "" {
		opts["query"] = query
	}
	if limit := c.Request.Form.Get("limit"); limit != "" {
		if limit, err := strconv.Atoi(limit); err == nil {
			opts["limit"] = limit
		}
	}
	if page := c.Request.Form.Get("page"); page != "" {
		if page, err := strconv.Atoi(page); err == nil {
			opts["page"] = page
		}
	}
	for _, param := range params {
		if v := c.Request.Form.Get(param); v != "" {
			if opts[param] == nil {
				opts[param] = v
			}
		}
	}
	return &opts
}
