package models

import "time"

type Comment struct {
	Id_RENAME int       `orm:"column(id)"`
	Uid       uint      `orm:"column(uid)" description:"用户ID"`
	Pid       uint      `orm:"column(pid)" description:"文章ID"`
	Reid      uint      `orm:"column(reid)" description:"上个回复ID"`
	Username  string    `orm:"column(username);size(50)" description:"姓名"`
	Usermail  string    `orm:"column(usermail);size(30)" description:"邮件地址"`
	Userurl   string    `orm:"column(userurl);size(50)" description:"链接地址"`
	Tipme     int8      `orm:"column(tipme)" description:"回复是否通知我"`
	Content   string    `orm:"column(content)" description:"留言内容"`
	Posttime  time.Time `orm:"column(posttime);type(datetime)" description:"评论时间"`
	Ip        string    `orm:"column(ip);size(20)" description:"IP地址"`
	Ishidden  int8      `orm:"column(ishidden)" description:"是否隐藏"`
	Ispass    int8      `orm:"column(ispass)" description:"是否通过"`
}
