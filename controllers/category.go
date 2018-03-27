package controllers

import (
	"fmt"
	"github.com/linauror/owncms-go/models"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Index() {
	slug := c.Ctx.Input.Param(":slug")
	categoryInfo, _ := models.GetCategoryBySlug(slug)
	fmt.Printf("%+v", categoryInfo)

	// 文章列表
	filter := make(map[string]string)
	orderBy := make([]string, 0)
	filter["category_slug"] = slug
	page, _ := c.GetInt64("page", 1)
	limit := int64(4)
	postLists, postTotal := models.GetAllPost(page, limit, orderBy, filter)

	c.Data["postTotal"] = postTotal
	c.Data["postLists"] = postLists
	c.Data["pagination"] = c.pagination(page, limit, postTotal)
	c.Data["categoryInfo"] = categoryInfo
	c.display()
}
