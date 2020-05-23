package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myfirst/models"
	"path"
	"time"
)

type IndexController struct {
	beego.Controller
}

//show Index
func (c *IndexController) Get() {
	o := orm.NewOrm()

	var articles []models.Article
	_, err := o.QueryTable("article").All(&articles)
	if err != nil {
		beego.Info(err)
	}
	c.Data["articles"] = articles

	c.TplName = "index.html"
}

//show add article
func (c *IndexController) ShowAddArticle() {
	c.TplName = "add.html"
}

//handle add article
func (c *IndexController) AddArticle() {
	//get data
	AName := c.GetString("articleName")
	AContent := c.GetString("content")
	//get image
	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	//limit format
	fileExt := path.Ext(h.Filename)
	beego.Info(fileExt)
	if fileExt != ".jpg" && fileExt != "png" {
		beego.Info("format not support")
		return
	}
	//limit size
	if h.Size > 5000000 {
		beego.Info("size")
		return
	}
	//unique img name
	fileName := time.Now().Format("2006-01-02-15-04-05") + fileExt
	//
	if err != nil {
		beego.Info("upload failed")
	} else {
		c.SaveToFile("uploadname", "./static/img/"+fileName)
		beego.Info(AName, AContent, fileName)

	}

	//validate format
	if AName == "" || AContent == "" {
		beego.Info("cant be nil")
	}

	//ORM
	o := orm.NewOrm()
	Article := models.Article{}
	Article.AName = AName
	Article.AContent = AContent
	Article.AImg = "/static/img/" + fileName
	_, err = o.Insert(&Article)
	if err != nil {
		beego.Info(err)
	}
	c.Redirect("/AddArticle", 302)
}
