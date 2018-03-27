package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Uid        int       `orm:"column(uid);PK"`
	Username   string    `orm:"column(username);size(20)" description:"用户名"`
	Password   string    `orm:"column(password);size(32)" description:"密码"`
	Salt       string    `orm:"column(salt);size(6)" description:"加密字符串"`
	Usermail   string    `orm:"column(usermail);size(20)" description:"用户邮箱"`
	Userurl    string    `orm:"column(userurl);size(50)" description:"用户网址"`
	Logintime  time.Time `orm:"column(logintime);type(datetime)" description:"登录时间"`
	Loginip    string    `orm:"column(loginip);size(19)" description:"登录IP"`
	Logedtime  time.Time `orm:"column(logedtime);type(datetime)" description:"上次登录时间"`
	Logedip    string    `orm:"column(logedip);size(19)" description:"上次登录IP"`
	Regip      string    `orm:"column(regip);size(19)" description:"注册IP"`
	Regtime    time.Time `orm:"column(regtime);type(datetime)" description:"注册时间"`
	Group      int8      `orm:"column(group)" description:"1管理员2编辑3普通会员"`
	Isverify   int8      `orm:"column(isverify)" description:"是否已经通过验证"`
	Status     int8      `orm:"column(status)" description:"1:正常0:禁止"`
	Logincount uint      `orm:"column(logincount)" description:"登录次数"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserByUsername(username string) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{Username: username}
	err = o.Read(user, "username")
	if err != nil {
		return nil, err
	}

	return user, nil
}
