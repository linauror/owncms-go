package main

import (
	"github.com/astaxie/beego"
	"github.com/linauror/owncms-go/controllers"
	_ "github.com/linauror/owncms-go/routers"
)

func main() {
	beego.ErrorController(&controllers.BaseController{})
	beego.Run()
}
