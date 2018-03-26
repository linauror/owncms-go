package controllers

import (
	"strconv"

	"github.com/linauror/owncms-go/models"
)

type PostController struct {
	BaseController
}

func (c *PostController) Info() {
	idstr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	post, _ := models.PostInfo(id)

	c.Data["post"] = post
	c.display()
}
