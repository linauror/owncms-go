package models

type Tag struct {
	Id_RENAME int    `orm:"column(id)"`
	Tag       string `orm:"column(tag);size(50)" description:"标签"`
}
