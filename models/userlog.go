package models

import "time"

type Userlog struct {
	Id_RENAME int       `orm:"column(id)"`
	Uid       uint      `orm:"column(uid)"`
	Msg       string    `orm:"column(msg);size(600)"`
	Ip        string    `orm:"column(ip);size(19)"`
	Time      time.Time `orm:"column(time);type(datetime)"`
	Type      int8      `orm:"column(type)" description:"1：后台操作2：用户操作"`
}
