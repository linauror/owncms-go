package admincontrollers

type PostController struct {
	BaseController
}

func (c *PostController) Index() {
	c.TplName = "admin/post.html"
}
