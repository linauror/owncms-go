package controllers

import (
	"github.com/astaxie/beego"

	"github.com/linauror/owncms-go/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {
	if c.Ctx.Input.IsGet() {
		c.Data["referer"] = c.Ctx.Request.Referer()
	} else {
		username := c.GetString("username")
		password := c.GetString("password")
		referer := c.GetString("referer")
		keepLogin, _ := c.GetBool("keepLogin")
		_, token, expiredTime, err := models.UserLogin(username, password, keepLogin)
		if err != nil {
			c.ShowError(err.Error())
			return
		}
		c.SetSecureCookie(beego.AppConfig.String("appkey"), "token", token, expiredTime)
		c.Redirect(referer, 301)
	}
	c.display()
}

func (c *UserController) Logout() {
	referer := c.Ctx.Request.Referer()
	c.SetSecureCookie(beego.AppConfig.String("appkey"), "token", "", -1)
	c.Redirect(referer, 301)
}

func (c *UserController) Register() {
	c.display()
}

func (c *UserController) Profile() {
	c.display()
}
