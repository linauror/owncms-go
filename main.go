package main

import (
	"github.com/astaxie/beego"
	"github.com/linauror/owncms-go/controllers"
	"github.com/linauror/owncms-go/library"
	_ "github.com/linauror/owncms-go/routers"
)

func main() {
	beego.ErrorController(&controllers.BaseController{})
	beego.AddFuncMap("Add", library.Add)
	beego.Run()
}
