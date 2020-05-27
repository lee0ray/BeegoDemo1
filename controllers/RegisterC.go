package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myfirst/models"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "Register.html"
}

func (c *RegisterController) Post() {
	//get info
	UserName := c.GetString("UserName")
	Pwd := c.GetString("Password")
	beego.Info(UserName, Pwd)
	if UserName == "" || Pwd == "" {
		beego.Info("data cannot be nil")
		c.Redirect("/Register", 302)
		return
	}
	//insert data
	o := orm.NewOrm()
	u := models.User{}
	u.Name = UserName
	u.Password = Pwd
	err := o.Read(&u, "Name")
	if err == nil {
		beego.Info("user exist")
		c.TplName = "Register.html"
	} else {
		_, err = o.Insert(&u)
		if err != nil {
			beego.Info("insert failed", err)
			c.Redirect("/Register", 302)
			return
		}
		c.Ctx.WriteString("Register Succeed")
		c.TplName = "Register.html"
	}
}

func (c *RegisterController) Login() {
	userName := c.Ctx.GetCookie("userName")
	if userName == ""{
		beego.Info("no cookie")
	}else {
		c.Data["userName"] = userName
	}
	c.TplName = "Login.html"
}

func (c *RegisterController) LoginPost() {
	UserName := c.GetString("UserName")
	Pwd := c.GetString("Password")
	if UserName == "" || Pwd == "" {
		beego.Info("data cannot be nil")
		c.Redirect("/Login", 302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Name = UserName
	user.Password = Pwd
	err := o.Read(&user, "name")
	if err != nil {
		beego.Info("no user")
		c.Ctx.WriteString("wrong user")
		c.Redirect("/Login", 302)
	}
	err2 := o.Read(&user, "password")
	if err2 != nil {
		beego.Info("password wrong")
		c.Ctx.WriteString("wrong password")
		c.Redirect("/Login", 302)
	}
	// check selection of rememeber me
	remember:=c.GetString("remember")
	beego.Info(remember)
	if remember =="on"{
		c.Ctx.SetCookie("userName",UserName,time.Second*3600)
	}else {
		c.Ctx.SetCookie("userName",UserName,-1)
	}

	c.SetSession("userName",UserName)
	c.Redirect("/Index", 302)
}
