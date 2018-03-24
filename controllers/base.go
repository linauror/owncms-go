package controllers

import (
	"fmt"
	"strings"

	"github.com/linauror/owncms-go/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (c *BaseController) Prepare() {
	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	c.actionName = strings.ToLower(actionName)
	menus, _ := models.MenuLists()
	fmt.Println(menus)
	c.Data["menus"] = menus
	c.Layout = "layout.html"
}

// Error404 404页面
func (c *BaseController) Error404() {
	c.Layout = ""
	c.TplName = "404.html"
}

func (c *BaseController) MsgTip(msg string) {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	flash.Error(msg)
	flash.Store(&c.Controller)
}

//加载模板
func (c *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = c.controllerName + "/" + c.actionName + ".html"
	}
	c.Layout = "layout.html"
	c.TplName = tplname
}
