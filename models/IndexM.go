package models

import "time"

type Article struct {
	Id          int       `orm:"pk;auto"`
	AName       string    `orm:"size(20);unique"`
	ATime       time.Time `orm:"column(a_time);type(datetime);auto_now_add;default(0)"`
	ACount      int       `orm:"default(0);null"`
	AContent    string
	AImg        string
	ArticleType *ArticleType `orm:"rel(fk)"`
	Users       []*User      `orm:"reverse(many)"`
}

type ArticleType struct {
	Id       int
	TypeName string     `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}
