package models

import "time"

type Page struct {
	Id_RENAME  int       `orm:"column(id)"`
	Title      string    `orm:"column(title);size(100)" description:"标题"`
	Slug       string    `orm:"column(slug);size(100)" description:"缩略标题"`
	Content    string    `orm:"column(content)" description:"内容"`
	Posttime   time.Time `orm:"column(posttime);type(datetime)" description:"发表时间"`
	Modifytime time.Time `orm:"column(modifytime);type(datetime)" description:"最后修改时间"`
	Uid        uint      `orm:"column(uid)" description:"作者ID"`
	Template   string    `orm:"column(template);size(50)" description:"模板名称"`
	Ishidden   int8      `orm:"column(ishidden)" description:"是否隐藏"`
}
