package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
)

// TODO: Move to middleware
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
		params["limit"] = limit
	}
	if limit, offset, ok := pagination(c); ok {
		params["limit"] = limit
		params["offset"] = offset
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

func pagination(c *gin.Context) (int, int, bool) {
	var (
		limit  int
		offset int
		ok     bool
	)
	if rangeData := c.Request.Header.Get("Range"); rangeData != "" {
		pattern := regexp.MustCompile("(?:items )?(\\d+)\\-(\\d+)")
		if match := pattern.FindStringSubmatch(rangeData); match != nil {
			if v, err := strconv.Atoi(match[1]); err == nil {
				offset = v
			}
			if v, err := strconv.Atoi(match[2]); err == nil {
				limit = v + 1 - offset
			}
			ok = true
		}
	}
	return limit, offset, ok
}

func writeRangeHeader(c *gin.Context, count int, limit int, offset int) {
	contentRangeHeader := fmt.Sprintf("*/0")
	if count > 0 {
		contentRangeHeader = fmt.Sprintf("%v-%v/%v", offset, offset+limit-1, count)

	}
	c.Writer.Header().Set("Accept-Ranges", "items")
	c.Writer.Header().Set("Range-Unit", "items")
	c.Writer.Header().Set("Content-Range", contentRangeHeader)
}
