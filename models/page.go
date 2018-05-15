package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Page struct {
	Id         int       `orm:"column(id)"`
	Title      string    `orm:"column(title);size(100)" description:"标题"`
	Slug       string    `orm:"column(slug);size(100)" description:"缩略标题"`
	Content    string    `orm:"column(content)" description:"内容"`
	Posttime   time.Time `orm:"column(posttime);type(datetime)" description:"发表时间"`
	Modifytime time.Time `orm:"column(modifytime);type(datetime)" description:"最后修改时间"`
	Template   string    `orm:"column(template);size(50)" description:"模板名称"`
	Ishidden   int8      `orm:"column(ishidden)" description:"是否隐藏"`
	User       *User     `orm:"rel(fk);column(uid)"`
}

func (t *Page) TableName() string {
	return "page"
}

func init() {
	orm.RegisterModel(new(Page))
}

func GetPageBySlug(slug string) (page *Page, err error) {
	o := orm.NewOrm()
	page = &Page{Slug: slug}
	err = o.Read(page, "slug")
	if err != nil {
		return nil, err
	}

	return page, nil
}

func GetAllPage(page, limit int64, orderBy []string, filter map[string]string) (lists []*Page, total int64) {
	offset := (page - 1) * limit
	query := orm.NewOrm().QueryTable(new(Page))
	if len(filter) > 0 {
		for k, v := range filter {
			switch k {
			case "title":
				query = query.Filter("title__contains", v)
			default:
				query = query.Filter(k, v)
			}
		}
	}

	total, _ = query.Count()

	if len(orderBy) > 0 {
		query = query.OrderBy(orderBy...)
	} else {
		query = query.OrderBy("-id")
	}

	query.Limit(limit, offset).RelatedSel().All(&lists)

	return lists, total
}
