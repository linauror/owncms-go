package controllers

import (
	"github.com/linauror/owncms-go/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	filter := make(map[string]string)

	title := c.GetString("title")
	if len(title) > 0 {
		filter["title"] = title
	}

	postLists, postTotal := models.PostLists(1, 3, filter)

	c.Data["postTotal"] = postTotal
	c.Data["postLists"] = postLists
	c.display()
}
