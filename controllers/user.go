package controllers

import (
	"strconv"

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
		c.Data["referer"] = c.Ctx.Input.Referer()
	} else {
		username := c.GetString("username")
		password := c.GetString("password")
		referer := c.GetString("referer")
		ip := c.Ctx.Input.IP()
		keepLogin, _ := c.GetBool("keepLogin")
		_, token, expiredTime, err := models.UserLogin(username, password, ip, keepLogin)
		if err != nil {
			c.ShowTip(err.Error())
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

		u.Regip = c.Ctx.Input.IP()

		_, err := models.UserRegister(&u)

		if err != nil {
			c.ShowTip(err.Error())
			return
		}

		c.ShowTip("恭喜您，注册成功，请登录")
		return
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

	if c.Ctx.Input.IsPost() {
		data := make(map[string]string)
		if len(c.GetString("usermail")) > 0 {
			data["usermail"] = c.GetString("usermail")
		}
		if len(c.GetString("userurl")) > 0 {
			data["userurl"] = c.GetString("userurl")
		}
		if len(c.GetString("password")) > 0 {
			data["password"] = c.GetString("password")
		}

		err := models.UserUpdate(c.AuthUser, data)
		if err != nil {
			c.ShowTip(err.Error())
			return
		}

		if _, exist := data["password"]; exist {
			c.SetSecureCookie(beego.AppConfig.String("appkey"), "token", "", -1)
			c.ShowTip("密码已更新，请重新登录")
			return
		}
	}

	c.Data["groupDesc"] = GROUPS[c.AuthUser.Group]
	c.display()
}
