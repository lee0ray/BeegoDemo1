package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	Articles []*Article `orm:"rel(m2m)"`
}

type Test struct {
	Id   int
	Time time.Time `orm:"auto_now"`
}

func init() {
	// set database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8")
	//mapping model
	orm.RegisterModel(new(Test))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Article))
	orm.RegisterModel(new(ArticleType))
	//generate table
	orm.RunSyncdb("default", false, true)
}

func TestCRUD() {
	o := orm.NewOrm()
	qs := o.QueryTable("demo")
	qs.Filter("id", 1)

}
