package admincontrollers

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Index() {
	c.TplName = "admin/category.html"
}
