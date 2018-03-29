package controllers

import (
	"strconv"

	"github.com/linauror/owncms-go/models"
)

type AuthorController struct {
	BaseController
}

func (c *AuthorController) Index() {
	username := c.Ctx.Input.Param(":username")
	userInfo, _ := models.GetUserByUsername(username)

	// 文章列表
	filter := make(map[string]string)
	orderBy := make([]string, 0)
	filter["uid"] = strconv.Itoa(userInfo.Uid)
	page, _ := c.GetInt64("page", 1)
	limit := int64(20)
	postLists, postTotal := models.GetAllPost(page, limit, orderBy, filter)

	c.Data["postTotal"] = postTotal
	c.Data["postLists"] = postLists
	c.Data["pagination"] = c.pagination(page, limit, postTotal)
	c.Data["userInfo"] = userInfo
	c.TplName = "author.html"
}
