package models

type Media struct {
	Id_RENAME int    `orm:"column(id)"`
	Uid       uint   `orm:"column(uid)"`
	Title     string `orm:"column(title);size(100)" description:"文件名称"`
	Des       string `orm:"column(des);size(200)" description:"文件描述"`
	Filepath  string `orm:"column(filepath);size(100)" description:"文件路径"`
	Suffix    string `orm:"column(suffix);size(10)" description:"文件后缀"`
}
