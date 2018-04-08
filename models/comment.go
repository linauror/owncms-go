package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id          int       `orm:"column(id)"`
	Uid         int       `orm:"column(uid)" description:"用户ID"`
	Pid         int       `orm:"column(pid)" description:"文章ID"`
	Reid        int       `orm:"column(reid)" description:"上个回复ID"`
	Username    string    `orm:"column(username);size(50)" description:"姓名"`
	Usermail    string    `orm:"column(usermail);size(30)" description:"邮件地址"`
	Userurl     string    `orm:"column(userurl);size(50)" description:"链接地址"`
	Tipme       int8      `orm:"column(tipme)" description:"回复是否通知我"`
	Content     string    `orm:"column(content)" description:"留言内容"`
	Posttime    time.Time `orm:"column(posttime);type(datetime)" description:"评论时间"`
	Ip          string    `orm:"column(ip);size(20)" description:"IP地址"`
	Ishidden    int8      `orm:"column(ishidden)" description:"是否隐藏"`
	Ispass      int8      `orm:"column(ispass)" description:"是否通过"`
	Slug        string    `orm:"column(slug)"`
	Title       string    `orm:"column(title)"`
	Channeltype string    `orm:"column(channeltype)"`
}

func (t *Comment) TableName() string {
	return "comment"
}

func init() {
	orm.RegisterModel(new(Comment))
}

func GetAllComment(page, limit int64, filter map[string]string) (lists []*Comment, total int64) {
	offset := (page - 1) * limit
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("comment.*,post.slug,post.title,category.channeltype").From("comment").LeftJoin("post").On("comment.pid = post.id").LeftJoin("category").On("post.category = category.id")

	query := orm.NewOrm().QueryTable(new(Comment))

	var where []string
	var params []string
	if len(filter) > 0 {
		for k, v := range filter {
			switch k {
			default:
				query = query.Filter(k, v)
				where = append(where, "comment."+k+" = ?")
				params = append(params, v)
			}
		}
	}
	if len(where) > 0 {
		qb.Where(strings.Join(where, " AND "))
	}

	total, _ = query.Count()

	qb.OrderBy("post.id").Desc().Limit(int(limit)).Offset(int(offset))

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, params).QueryRows(&lists)

	return lists, total
}

func (comment *Comment) Pub() (err error) {
	_, err = orm.NewOrm().Insert(comment)
	return err
}
