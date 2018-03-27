package controllers

import (
	"github.com/linauror/owncms-go/models"
)

type PageController struct {
	BaseController
}

func (c *PageController) Info() {
	slug := c.Ctx.Input.Param(":slug")
	PageInfo, _ := models.GetPageBySlug(slug)

	c.Data["page"] = PageInfo
	c.display()
}
