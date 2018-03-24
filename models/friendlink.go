package models

type Friendlink struct {
	Id_RENAME int    `orm:"column(id)"`
	Logo      string `orm:"column(logo);size(100)" description:"略缩图"`
	Title     string `orm:"column(title);size(100)" description:"标题"`
	Dec       string `orm:"column(dec);size(100)" description:"链接描述"`
	Url       string `orm:"column(url);size(100)" description:"链接地址"`
	Sortrank  uint   `orm:"column(sortrank)" description:"排序 倒序"`
	Ishidden  int8   `orm:"column(ishidden)" description:"是否隐藏"`
}
