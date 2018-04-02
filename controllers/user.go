package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"github.com/linauror/owncms-go/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {
	if c.Uid != 0 {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}

	if c.Ctx.Input.IsGet() {
		c.Data["referer"] = c.Ctx.Request.Referer()
	} else {
		username := c.GetString("username")
		password := c.GetString("password")
		referer := c.GetString("referer")
		ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")
		keepLogin, _ := c.GetBool("keepLogin")
		_, token, expiredTime, err := models.UserLogin(username, password, ip[0], keepLogin)
		if err != nil {
			c.ShowError(err.Error())
			return
		}
		c.SetSecureCookie(beego.AppConfig.String("appkey"), "token", token, expiredTime)
		c.Redirect(referer, 302)
	}
	c.display()
}

func (c *UserController) Logout() {
	referer := c.Ctx.Request.Referer()
	c.SetSecureCookie(beego.AppConfig.String("appkey"), "token", "", -1)
	c.Redirect(referer, 302)
}

func (c *UserController) Register() {
	if c.Ctx.Input.IsPost() {
		u := models.User{}
		c.ParseForm(&u)

		ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")
		u.Regip = ip[0]

		_, err := models.UserRegister(&u)

		if err != nil {
			c.ShowError(err.Error())
			return
		}

		c.Redirect("/", 302)
	}
	c.display()
}

func (c *UserController) Profile() {
	c.CheckLogin()

	if c.AuthUser.Group <= 2 {
		filter := make(map[string]string)
		orderBy := make([]string, 0)
		filter["uid"] = strconv.Itoa(c.AuthUser.Uid)
		_, postTotal := models.GetAllPost(int64(1), int64(1), orderBy, filter)
		c.Data["postTotal"] = postTotal
	}

	c.Data["groupDesc"] = GROUPS[c.AuthUser.Group]
	c.display()
}
