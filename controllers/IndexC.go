package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"myfirst/models"
	"path"
	"strconv"
	"time"
)

type IndexController struct {
	beego.Controller
}

//show Index
func (c *IndexController) Get() {
	o := orm.NewOrm()

	var articles []models.Article
	qs:= o.QueryTable("article")
	//qs.All(&articles)
	count,err := qs.Count()
	if err!=nil{
		beego.Info(err)
	}
	//set page
	var PageIndex1 int
	PageIndex :=c.GetString("pageIndex")
	PageIndex1,err =strconv.Atoi(PageIndex)
	if err!=nil{
		PageIndex1 =1
		beego.Info(err)
	}

	beego.Info(count)
	PageSize:= 3
	start:= PageSize*(PageIndex1-1)
	pageCount := math.Ceil(float64(count)/float64(PageSize))
	qs.Limit(PageSize,start).All(&articles)
	//
	c.Data["articles"] = articles
	c.Data["count"] = count
	c.Data["pagecount"] = pageCount
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
	c.Redirect("/Index", 302)
}

func (c *IndexController) HandleDelete() {
	Id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
	}

	o := orm.NewOrm()
	Article := models.Article{Id: Id}
	err = o.Read(&Article)
	if err != nil {
		beego.Info(err)
	} else {
		o.Delete(&Article)
		c.Redirect("/Index",302)
	}
}
