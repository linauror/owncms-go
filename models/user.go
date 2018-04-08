package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/linauror/owncms-go/library"

	"github.com/astaxie/beego/orm"
)

const expireTime = 86400 * 30

type User struct {
	Uid        int       `orm:"column(uid);PK"`
	Username   string    `orm:"column(username);size(20)" description:"用户名" form:"username"`
	Password   string    `orm:"column(password);size(32)" description:"密码" form:"password"`
	Salt       string    `orm:"column(salt);size(6)" description:"加密字符串"`
	Usermail   string    `orm:"column(usermail);size(20)" description:"用户邮箱" form:"usermail"`
	Userurl    string    `orm:"column(userurl);size(50)" description:"用户网址" form:"userurl"`
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
		return nil, errors.New("不存在的用户")
	}

	return user, nil
}

func GetUserByUsermail(usermail string) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{Usermail: usermail}
	err = o.Read(user, "usermail")
	if err != nil {
		return nil, errors.New("不存在的用户")
	}

	return user, nil
}

func UserLogin(username, password string, ip string, keepLogin bool) (user *User, token string, expiredTime int, err error) {
	user, err = GetUserByUsername(username)
	if err != nil {
		return nil, "", 0, err
	}

	str := password + user.Salt
	h := md5.New()
	io.WriteString(h, str)
	hash := h.Sum(nil)

	if user.Password != hex.EncodeToString(hash) {
		return nil, "", 0, errors.New("密码错误")
	} else if user.Status != 1 {
		return nil, "", 0, errors.New("该用户已经被禁用")
	}

	now := int(time.Now().Unix())
	expiredTime = 0
	if keepLogin {
		expiredTime = now + expireTime
	}
	str = user.Username + "|" + strconv.Itoa(expiredTime)
	h = md5.New()
	io.WriteString(h, str+"|"+user.Password)
	hash = h.Sum(nil)
	token = str + "|" + string(hash)

	user.Logedip = user.Loginip
	user.Logedtime = user.Logintime
	user.Loginip = ip
	user.Logintime = time.Now()
	user.Logincount = user.Logincount + 1
	orm.NewOrm().Update(user, "logedip", "logedtime", "loginip", "logintime", "logincount")

	return user, token, expiredTime, nil
}

func UserRegister(u *User) (user *User, err error) {
	checkuser, err := GetUserByUsername(u.Username)
	if checkuser != nil {
		return nil, errors.New("用户名已存在")
	}

	checkuser, err = GetUserByUsername(u.Usermail)
	if checkuser != nil {
		return nil, errors.New("用户邮箱已存在")
	}

	salt := library.GetRandomString(6)
	str := u.Password + salt
	h := md5.New()
	io.WriteString(h, str)
	hash := h.Sum(nil)

	u.Salt = salt
	u.Password = hex.EncodeToString(hash)
	u.Regtime = time.Now()
	u.Group = 3
	u.Isverify = 0
	u.Status = 1
	u.Logedip = u.Regip
	u.Loginip = u.Regip
	u.Logedtime = time.Now()
	u.Logintime = time.Now()

	_, err = orm.NewOrm().Insert(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func UserAuth(token string) (user *User, err error) {
	str := strings.Split(token, "|")
	if len(str) != 3 {
		return nil, errors.New("token无效")
	}

	user, err = GetUserByUsername(str[0])
	if err != nil {
		return nil, err
	}

	expiredTime, _ := strconv.Atoi(str[1])
	now := int(time.Now().Unix())
	if now <= expiredTime {
		return nil, errors.New("token已过期")
	}

	hashstr := user.Username + "|" + str[1]
	h := md5.New()
	io.WriteString(h, hashstr+"|"+user.Password)
	hash := h.Sum(nil)
	if string(hash) != str[2] {
		return nil, errors.New("token无效")
	}

	return user, nil
}

func UserUpdate(user *User, data map[string]string) (err error) {
	var fields []string
	for k, v := range data {
		fields = append(fields, k)
		switch k {
		case "usermail":
			user.Usermail = v
		case "userurl":
			user.Userurl = v
		case "password":
			salt := library.GetRandomString(6)
			str := v + salt
			h := md5.New()
			io.WriteString(h, str)
			hash := h.Sum(nil)
			user.Salt = salt
			user.Password = hex.EncodeToString(hash)
			fields = append(fields, "salt")
		}
	}

	_, err = orm.NewOrm().Update(user, fields...)
	return err
}
