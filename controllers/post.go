package controllers

import (
	"strconv"
	"strings"

	"github.com/linauror/owncms-go/models"
)

type PostController struct {
	BaseController
}

func (c *PostController) Info() {
	idstr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	post, _ := models.GetPostById(id)

	if _tagIds := post.Tag; len(_tagIds) > 0 {
		tagIds := strings.Split(_tagIds, ",")
		tags := models.GetTagsByIds(tagIds)
		c.Data["tags"] = tags
	}

	nextPost, prevPost := models.GetNextAndProvPost(id)

	commentFilter := make(map[string]string)
	commentFilter["pid"] = idstr
	commentLists, _ := models.GetAllComment(int64(1), int64(1000), commentFilter)

	c.Data["nextPost"] = nextPost
	c.Data["prevPost"] = prevPost
	c.Data["commentLists"] = commentLists
	c.Data["post"] = post
	c.TplName = "post.html"
}
