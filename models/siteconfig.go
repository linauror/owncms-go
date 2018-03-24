package models

import "github.com/astaxie/beego/orm"

type Siteconfig struct {
	Id          int8   `orm:"column(id);PK"`
	Varname     string `orm:"column(varname);size(50)" description:"变量名称"`
	Description string `orm:"column(description);size(50)" description:"变量描述"`
	Value       string `orm:"column(value)" description:"变量值"`
	Inputtype   string `orm:"column(inputtype);size(10)" description:"值类型"`
}

func (t *Siteconfig) TableName() string {
	return "sietconfig"
}

func init() {
	orm.RegisterModel(new(Siteconfig))
}

func SiteconfigLists() ([]*Siteconfig, int64) {
	query := orm.NewOrm().QueryTable(new(Siteconfig))
	total, _ := query.Count()
	list := make([]*Siteconfig, 0)
	query.OrderBy("-id").All(&list)

	return list, total
}
