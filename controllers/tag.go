package controllers

import (
	"github.com/astaxie/beego"
	"github.com/linauror/owncms-go/models"
)

type TagController struct {
	BaseController
}

func (c *TagController) Index() {
	tag := c.Ctx.Input.Param(":tag")
	tagInfo, _ := models.GetTagByStr(tag)

	beego.Debug("tag:" + tag)

	page, _ := c.GetInt64("page", 1)
	limit := int64(20)
	postLists, postTotal := models.GetPostsByTag(tagInfo.Id, page, limit)

	c.Data["postTotal"] = postTotal
	c.Data["postLists"] = postLists
	c.Data["pagination"] = c.pagination(page, limit, postTotal)
	c.Data["tag"] = tag
	c.TplName = "tag.html"
}
