package controllers

import (
	"github.com/linauror/owncms-go/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	filter := make(map[string]string)
	postLists, postTotal := models.PostLists(1, 10, filter)

	c.Data["postTotal"] = postTotal
	c.Data["postLists"] = postLists
	c.display()
}
