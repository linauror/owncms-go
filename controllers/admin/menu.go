package admincontrollers

type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {
	c.TplName = "admin/menu.html"
}
