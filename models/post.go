package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id            int       `orm:"column(id)"`
	CategoryId    uint8     `orm:"column(category)" description:"类别ID"`
	Flag          string    `orm:"column(flag);size(50)" description:"属性"`
	Litpic        string    `orm:"column(litpic);size(100)" description:"略缩图"`
	Title         string    `orm:"column(title);size(100)" description:"标题"`
	Slug          string    `orm:"column(slug);size(100)" description:"缩略标题"`
	Content       string    `orm:"column(content)" description:"内容"`
	Keyword       string    `orm:"column(keyword);size(200)" description:"关键词"`
	Description   string    `orm:"column(description);size(600)" description:"描述"`
	Source        string    `orm:"column(source);size(50)" description:"文章来源"`
	Click         uint      `orm:"column(click)" description:"点击次数"`
	CommentCount  uint      `orm:"column(comment_count)" description:"评论次数"`
	CommentStatus int8      `orm:"column(comment_status)" description:"是否允许评论"`
	Uid           uint      `orm:"column(uid)" description:"作者ID"`
	Sortrank      uint      `orm:"column(sortrank)" description:"排序 倒序"`
	Posttime      time.Time `orm:"column(posttime);type(datetime)" description:"发表时间"`
	Modifytime    time.Time `orm:"column(modifytime);type(datetime)" description:"最后修改时间"`
	Template      string    `orm:"column(template);size(50)" description:"模板名称"`
	Tag           string    `orm:"column(tag);size(100)" description:"标签ID集合"`
	Ishidden      int8      `orm:"column(ishidden)" description:"是否隐藏"`
	Channeltype   string
}

func (t *Post) TableName() string {
	return "post"
}

func init() {
	orm.RegisterModel(new(Post))
}

func PostLists(page int, pageSize int, filter map[string]string) ([]*Post, int64) {
	offset := (page - 1) * pageSize
	qb, _ := orm.NewQueryBuilder("mysql")
	list := make([]*Post, 0)
	qb.Select("post.*", "category.channeltype").From("post").LeftJoin("category").On("post.category = category.id")

	var where []string
	var cond []string
	if len(filter) > 0 {
		for k, v := range filter {
			switch k {
			case "title":
				where = append(where, "post.title like ?")
				cond = append(cond, "%"+v+"%")
				break
			default:
				where = append(where, "post."+k+" = ?")
				cond = append(cond, v)
				break
			}
		}
	}

	if len(where) > 0 {
		qb.Where(strings.Join(where, " AND "))
	}

	qb.OrderBy("post.id").Desc().Limit(pageSize).Offset(offset)

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, cond).QueryRows(&list)

	/*
		query := orm.NewOrm().QueryTable(new(Post))
		if len(filter) > 0 {
			for k, v := range filter {
				switch k {
				case "title":
					query.Filter("title__contains", v)
					break
				default:
					query.Filter(k, v)
					break
				}
			}
		}

		total, _ := query.Count()
		list := make([]*Post, 0)
		query.OrderBy("-sortrank", "-id").Limit(pageSize, offset).All(&list)
	*/

	return list, 10
}
