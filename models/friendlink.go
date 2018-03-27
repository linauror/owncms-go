package models

import "github.com/astaxie/beego/orm"

type Friendlink struct {
	Id       int    `orm:"column(id)"`
	Logo     string `orm:"column(logo);size(100)" description:"略缩图"`
	Title    string `orm:"column(title);size(100)" description:"标题"`
	Dec      string `orm:"column(dec);size(100)" description:"链接描述"`
	Url      string `orm:"column(url);size(100)" description:"链接地址"`
	Sortrank uint   `orm:"column(sortrank)" description:"排序 倒序"`
	Ishidden int8   `orm:"column(ishidden)" description:"是否隐藏"`
}

func (t *Friendlink) TableName() string {
	return "friendlink"
}

func init() {
	orm.RegisterModel(new(Friendlink))
}

func GetAllFriendlink(page, limit int64) (lists []*Friendlink, total int64) {
	offset := (page - 1) * limit
	query := orm.NewOrm().QueryTable(new(Friendlink))

	total, _ = query.Count()
	query.Limit(limit, offset).OrderBy("-id").All(&lists)

	return lists, total
}
