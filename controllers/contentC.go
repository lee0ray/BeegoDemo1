package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myfirst/models"
	"path"
	"time"
)

type ContentController struct {
	beego.Controller
}

//show content
func (c *ContentController) Get() {
	//get article id
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		return
	}
	beego.Info(id)
	// search database
	o := orm.NewOrm()
	Article := models.Article{Id: id}
	err = o.Read(&Article)
	if err != nil {
		beego.Info(err)
		return
	}
	c.Data["Article"] = Article
	c.TplName = "content.html"
}

//show update page
func (c *ContentController) ShowUpdate() {
	//get article id
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		return
	}
	beego.Info(id)
	// search database
	o := orm.NewOrm()
	Article := models.Article{Id: id}
	err = o.Read(&Article)
	if err != nil {
		beego.Info(err)
		return
	}
	c.Data["Article"] = Article
	c.TplName = "update.html"
}

//handle update
/*func (c *ContentController) Update() {
	//get data
	Id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("get id failed")
		return
	}
	AName := c.GetString("articleName")
	AContent := c.GetString("content")
	beego.Info(AName, AContent)
	//get image
	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	var fileName string
	//limit format
	if err != nil {
		beego.Info("upload failed")
		c.Ctx.WriteString("get img failed")
	} else {
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
		fileName = time.Now().Format("2006-01-02-15-04-05") + fileExt
		//
		c.SaveToFile("uploadname", "./static/img/"+fileName)
		beego.Info(AName, AContent, fileName)

	}
	//validate data
	if AName == "" || AContent == "" {
		beego.Info("update failed")
	}
	//update data
	o := orm.NewOrm()

	Article := models.Article{Id: Id}
	err = o.Read(&Article)
	if err != nil {
		beego.Info(err)
		return
	}
	Article.AName = AName
	Article.AContent = AContent
	Article.AImg = "./static/img/" + fileName
	_, err = o.Update(&Article, "AName", "AContent", "AImg")
	if err != nil {
		beego.Info(err)
		return
	}
	c.Redirect("/Index", 302)
	//
}*/
//handle update
func (c *ContentController) Update() {
	//get data
	Id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("get id failed")
		return
	}
	AName := c.GetString("articleName")
	AContent := c.GetString("content")
	beego.Info(AName, AContent)
	//get image
	f, h, err := c.GetFile("uploadname")
	if err !=nil{
		beego.Info(err)
		c.Redirect("/Index",302)
	}
	defer f.Close()
	var fileName string
	if err != nil {
		beego.Info("upload failed")
		c.Ctx.WriteString("get img failed")
	} else {
		//limit size
		//unique img name
		//
		//limit format
		fileExt := path.Ext(h.Filename)
		beego.Info(fileExt)
		if fileExt != ".jpg" && fileExt != "png" {
			beego.Info("format not support")
			return
		}
		if h.Size > 5000000 {
			beego.Info("size")
			return
		}
		fileName = time.Now().Format("2006-01-02-15-04-05") + fileExt
		c.SaveToFile("uploadname", "./static/img/"+fileName)
		beego.Info(AName, AContent, fileName)

	}

	//validate data
	if AName == "" || AContent == "" {
		beego.Info("update failed")
	}
	//update data
	o := orm.NewOrm()

	Article := models.Article{Id: Id}
	err = o.Read(&Article)
	if err != nil {
		beego.Info(err)
		return
	}
	Article.AName = AName
	Article.AContent = AContent
	Article.AImg = "./static/img/" + fileName
	beego.Info(AName, AContent, fileName)
	_, err = o.Update(&Article, "AName", "AContent", "AImg")
	if err != nil {
		beego.Info(err)
		return
	}
	c.Redirect("/Index", 302)
	//
}