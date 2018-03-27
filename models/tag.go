package models

import (
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id  int    `orm:"column(id)"`
	Tag string `orm:"column(tag);size(50)" description:"标签"`
}

func (t *Tag) TableName() string {
	return "tag"
}

func init() {
	orm.RegisterModel(new(Tag))
}

func GetTagsByIds(ids []string) (lists []*Tag) {
	orm.NewOrm().QueryTable(new(Tag)).Filter("id__in", ids).All(&lists)
	return lists
}
