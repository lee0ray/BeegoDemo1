package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myfirst/models"
)

type MainController struct {
	beego.Controller
}
type DemoController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *DemoController) Get() {

	//orm object
/*	o := orm.NewOrm()
	user := models.User{}
	user.Name = "user1"
	user.Password = "123"
	// insert data
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info(err)
		return
	}*/
	o :=orm.NewOrm()
	u := models.User{}
	for i:=3;i<100;i++ {
		u.Id = i
		if u.Name!="" {
			_, err := o.Delete(&u)
			if err != nil {
				beego.Info("delete failed", err)
			} else {
				beego.Info("delete completed", u.Id)
			}
		}
	}

	c.Data["data"] = "demo0"
	c.TplName = "demo0.html"
}

func (c *DemoController) Post() {
	o := orm.NewOrm()
	user := models.User{}
	 user.Id = 1
	 err:=o.Read(&user)
	if err!=nil {
		beego.Info("search failed",err)
	}else {
		beego.Info(user)
		user.Name = "user1new"
		user.Password="1234"

		_ ,err = o.Update(&user)
		if err!= nil{
			beego.Info("update failed",err)
		}else {
			beego.Info("update succeed",user)
		}
	}
	test:= models.Test{}
	_,err=o.Insert(&test)
	if err!=nil{
		beego.Info(err)
	}

	c.Data["data"] = "demo1"
	c.TplName = "demo0.html"
}

func (c *DemoController) Demo1Get() {
	o := orm.NewOrm()
	user := models.User{}
	user.Name = "user1"
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Info("search failed")
		return
	}

	beego.Info("search succeed\n","user info:",user)


	c.Data["data"] = "demo1"
	c.TplName = "demo1.html"
}
