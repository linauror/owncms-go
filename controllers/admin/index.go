package admincontrollers

import (
	"github.com/astaxie/beego"
	"github.com/linauror/owncms-go/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Login() {
	if c.Uid != 0 {
		c.Redirect("/admin/", 302)
	}

	if c.Ctx.Input.IsPost() {
		username := c.GetString("username")
		password := c.GetString("password")
		ip := c.Ctx.Input.IP()
		keepLogin, _ := c.GetBool("keepLogin")
		_, token, expiredTime, err := models.UserLogin(username, password, ip, keepLogin)
		if err != nil {
			c.ShowTip(err.Error())
			return
		}
		c.SetSecureCookie(beego.AppConfig.String("appkey"), "token", token, expiredTime)
		c.Redirect("admin/", 302)
	}
	c.Layout = ""
	c.TplName = "admin/login.html"
}

func (c *IndexController) Index() {
	c.TplName = "admin/index.html"
}
