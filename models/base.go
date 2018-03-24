package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpasswd := beego.AppConfig.String("db.passwd")
	dbcharset := beego.AppConfig.String("db.charset")
	dbtimezone := beego.AppConfig.String("db.timezone")
	dbname := beego.AppConfig.String("db.name")
	dsn := dbuser + ":" + dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + dbcharset + "&loc=" + url.QueryEscape(dbtimezone)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.Debug = true
}
