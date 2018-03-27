package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id          int    `orm:"column(id);auto"`
	Slug        string `orm:"column(slug);size(50)" description:"缩略标题"`
	Reid        int8   `orm:"column(reid)" description:"上层ID"`
	Typename    string `orm:"column(typename);size(50)" description:"类别名称"`
	Channeltype string `orm:"column(channeltype);size(50)" description:"频道模型"`
	Tempindex   string `orm:"column(tempindex);size(30)" description:"模板首页或者列表路径"`
	Temparticle string `orm:"column(temparticle);size(30)" description:"模板内容路径"`
}

func (t *Category) TableName() string {
	return "category"
}

func init() {
	orm.RegisterModel(new(Category))
}

func GetCategoryBySlug(slug string) (category *Category, err error) {
	o := orm.NewOrm()
	category = &Category{Slug: slug}
	err = o.Read(category, "slug")
	if err != nil {
		return nil, err
	}

	return category, nil
}
