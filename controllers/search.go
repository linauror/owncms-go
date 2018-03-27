package controllers

import (
	"github.com/linauror/owncms-go/models"
)

type SearchController struct {
	BaseController
}

func (c *SearchController) Index() {
	q := c.GetString("q")
	filter := make(map[string]string)
	orderBy := make([]string, 0)
	filter["title"] = q
	page, _ := c.GetInt64("page", 1)
	limit := int64(20)
	postLists, postTotal := models.GetAllPost(page, limit, orderBy, filter)

	c.Data["q"] = q
	c.Data["postTotal"] = postTotal
	c.Data["postLists"] = postLists
	c.Data["pagination"] = c.pagination(page, limit, postTotal)
	c.display()
}
