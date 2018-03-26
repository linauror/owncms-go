package controllers

import (
	"github.com/linauror/owncms-go/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	// 文章列表
	filter := make(map[string]string)
	var orderBy []string
	title := c.GetString("title")
	if len(title) > 0 {
		filter["title"] = title
	}
	page, _ := c.GetInt64("page", 1)
	limit := int64(4)
	postLists, postTotal := models.PostLists(page, limit, orderBy, filter)

	c.Data["postTotal"] = postTotal
	c.Data["postLists"] = postLists
	c.Data["pagination"] = c.pagination(page, limit, postTotal)
	c.display()
}
