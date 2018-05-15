package admincontrollers

type SiteconfigController struct {
	BaseController
}

func (c *SiteconfigController) Index() {
	c.TplName = "admin/siteconfig.html"
}
