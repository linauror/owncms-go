package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id            int       `orm:"column(id)"`
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
	Category      *Category `orm:"rel(fk);column(category)"`
}

type TagPost struct {
	Post
	Channeltype string `orm:"column(channeltype)"`
	Typename    string `orm:"column(typename)"`
}

func (t *Post) TableName() string {
	return "post"
}

func init() {
	orm.RegisterModel(new(Post))
}

func GetAllPost(page, limit int64, orderBy []string, filter map[string]string) (lists []*Post, total int64) {
	offset := (page - 1) * limit
	query := orm.NewOrm().QueryTable(new(Post))
	if len(filter) > 0 {
		for k, v := range filter {
			switch k {
			case "title":
				query = query.Filter("title__contains", v)
			case "category_slug":
				query = query.Filter("category__slug", v)
			default:
				query = query.Filter(k, v)
			}
		}
	}

	total, _ = query.Count()

	if len(orderBy) > 0 {
		query = query.OrderBy(orderBy...)
	} else {
		query = query.OrderBy("-sortrank", "-id")
	}

	query.Limit(limit, offset).RelatedSel().All(&lists)

	return lists, total
}

func GetPostById(id int) (p *Post, err error) {
	o := orm.NewOrm()
	p = &Post{Id: id}
	err = o.Read(p)
	if err != nil {
		return nil, err
	}
	o.LoadRelated(p, "Category")

	return p, nil
}

func GetNextAndProvPost(id int) (n Post, p Post) {
	orm.NewOrm().QueryTable(new(Post)).Filter("id__gt", id).RelatedSel().One(&n)
	orm.NewOrm().QueryTable(new(Post)).Filter("id__lt", id).OrderBy("-id").RelatedSel().One(&p)
	return n, p
}

func GetPostsByTag(tagId int, page, limit int64) (lists []*TagPost, total int64) {
	offset := (page - 1) * limit
	sql := "SELECT post.*,category.typename,category.channeltype FROM post LEFT JOIN category ON category.id = post.category WHERE FIND_IN_SET(?, tag) LIMIT ?,?"
	total, _ = orm.NewOrm().Raw(sql, tagId, offset, limit).QueryRows(&lists)

	return lists, total
}

func (post *Post) NewView() {
	post.Click = post.Click + 1
	orm.NewOrm().Update(post, "click")
}

func (post *Post) NewComment() {
	post.CommentCount = post.CommentCount + 1
	orm.NewOrm().Update(post, "comment_count")
}
