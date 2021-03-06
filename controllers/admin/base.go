package admincontrollers

import (
	"fmt"
	"math"
	"strings"

	"github.com/linauror/owncms-go/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var GROUPS = map[int8]string{1: "管理员", 2: "编辑", 3: "普通会员"}

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	AuthUser       *models.User
	Uid            int
	Username       string
	UserGroup      int8
}

func (c *BaseController) Prepare() {
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/log.log"}`)

	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	c.actionName = strings.ToLower(actionName)

	// 登录验证
	token, isExistToken := c.GetSecureCookie(beego.AppConfig.String("appkey"), "token")
	if isExistToken {
		user, err := models.UserAuth(token)
		if err == nil {
			c.AuthUser = user
			c.Uid = user.Uid
			c.Username = user.Username
			c.UserGroup = user.Group
		}
	}

	c.Data["Uid"] = c.Uid
	c.Data["AuthUser"] = c.AuthUser
	c.Data["controllerName"] = controllerName
	c.Layout = "admin/layout.html"
}

// Error404 404页面
func (c *BaseController) Error404() {
	c.Layout = ""
	c.Data["errorTitle"] = "页面没有找到"
	c.Data["error"] = "404"
	c.TplName = "404.html"
}

// 错误提示
func (c *BaseController) ShowTip(error string) {
	c.Layout = ""
	c.Data["errorTitle"] = "提示"
	c.Data["error"] = error
	c.TplName = "404.html"
}

func (c *BaseController) CheckLogin() {
	if c.Uid == 0 {
		c.Redirect("/admin/login", 302)
	}
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

// pagination 分页
func (c *BaseController) pagination(page, limit, total int64) string {
	if total == 0 || total <= limit {
		return ""
	}
	totalPage := int64(math.Ceil(float64(total) / float64(limit)))
	str := `<div class="pagination">`
	for i := int64(1); i <= totalPage; i++ {
		if page == i {
			str = str + fmt.Sprintf(`&nbsp;%d`, i)
		} else {
			str = str + fmt.Sprintf(`&nbsp;<a href="?page=%d">%d</a>`, i, i)
		}
	}
	str = str + `</div>`
	return str
}
