package admincontrollers

import "github.com/linauror/owncms-go/models"

type PageController struct {
	BaseController
}

func (c *PageController) Index() {
	filter := make(map[string]string)
	var orderBy []string
	title := c.GetString("title")
	if len(title) > 0 {
		filter["title"] = title
	}
	page, _ := c.GetInt64("page", 1)
	limit := int64(4)
	pageLists, pageTotal := models.GetAllPage(page, limit, orderBy, filter)

	c.Data["pageTotal"] = pageTotal
	c.Data["pageLists"] = pageLists
	c.Data["pagination"] = c.pagination(page, limit, pageTotal)
	c.TplName = "admin/page.html"
}
