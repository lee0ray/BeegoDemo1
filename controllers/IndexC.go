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
	//login session
	userName := c.GetSession("userName")
	if userName == nil {
		c.Redirect("/Login", 302)
	}
	//
	var articles []models.Article
	qs := o.QueryTable("article")
	//qs.All(&articles)
	count, err := qs.Count()

	if err != nil {
		beego.Info(err)
	}
	//set page
	var PageIndex1 int
	PageIndex := c.GetString("pageIndex")
	PageIndex1, err = strconv.Atoi(PageIndex)

	if err != nil {
		PageIndex1 = 1
		beego.Info(err)
	}
	beego.Info(count)
	PageSize := 3
	start := PageSize * (PageIndex1 - 1)
	pageCount := math.Ceil(float64(count) / float64(PageSize))
	qs.Limit(PageSize, start).RelatedSel("ArticleType").All(&articles)
	//set type
	var types []models.ArticleType
	qs1 := o.QueryTable("ArticleType")
	qs1.All(&types)

	beego.Info(types)
	c.Data["types"] = types
	//
	c.Data["articles"] = articles
	c.Data["count"] = count
	c.Data["pagecount"] = pageCount
	c.TplName = "index.html"
}

// show select type of articles
func (c *IndexController) Post() {
	typeName := c.GetString("select")
	beego.Info(typeName)
	if typeName == "" {
		beego.Info("submit type failed")
	}

	o := orm.NewOrm()

	var PageIndex1 int
	var types []models.ArticleType
	var articles []models.Article
	PageIndex := c.GetString("pageIndex")
	PageIndex1, err := strconv.Atoi(PageIndex)
	if err != nil {
		PageIndex1 = 1
		beego.Info(err)
	}
	//show articles based on selected type
	qs := o.QueryTable("Article")
	qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)
	count := len(articles)
	PageSize := 3
	start := PageSize * (PageIndex1 - 1)
	pageCount := math.Ceil(float64(count) / float64(PageSize))
	qs = o.QueryTable("Article")
	qs.Limit(PageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)
	beego.Info(articles)
	beego.Info(len(articles))
	qs1 := o.QueryTable("ArticleType")
	qs1.All(&types)
	beego.Info(types)
	//insert data in Html

	c.Data["types"] = types
	//
	c.Data["articles"] = articles
	c.Data["count"] = count
	c.Data["pagecount"] = pageCount
	c.TplName = "index.html"

}

//show add article
func (c *IndexController) ShowAddArticle() {
	//select type
	var types []models.ArticleType
	o := orm.NewOrm()
	qs1 := o.QueryTable("ArticleType")
	qs1.All(&types)
	c.Data["types"] = types
	c.TplName = "add.html"
}

//handle add article
func (c *IndexController) AddArticle() {
	//get data
	AName := c.GetString("articleName")
	AContent := c.GetString("content")
	//get image
	f, h, err := c.GetFile("uploadname")

	if err != nil {
		beego.Info(err)
		c.Redirect("/AddArticle", 302)

	}
	defer f.Close()
	//limit format
	fileExt := path.Ext(h.Filename)
	beego.Info(fileExt)
	if fileExt != ".jpg" && fileExt != ".png" {
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
	//
	//insert ArticleType
	typeName := c.GetString("select")
	if typeName == "" {
		beego.Info("get type failed")
	}
	var ArticleType models.ArticleType
	ArticleType.TypeName = typeName
	err = o.Read(&ArticleType, "TypeName")
	if err != nil {
		beego.Info(err)
	}
	Article.ArticleType = &ArticleType

	//
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
		c.Redirect("/Index", 302)
	}
}

func (c *IndexController) ShowAddType() {
	//read type
	o := orm.NewOrm()
	var ArticleType []models.ArticleType
	qs := o.QueryTable("ArticleType")
	_, err := qs.All(&ArticleType)
	if err != nil {
		beego.Info(err)
	}
	c.Data["types"] = ArticleType
	c.TplName = "addType.html"
}

func (c *IndexController) AddType() {
	//get data
	o := orm.NewOrm()
	TypeName := c.GetString("TypeName")
	//validate data
	if TypeName == "" {
		beego.Info("type is null")
	}
	//insert data
	var ArticleType models.ArticleType
	ArticleType.TypeName = TypeName
	_, err := o.Insert(&ArticleType)
	if err != nil {
		beego.Info(err)
	}

	c.Redirect("/AddType", 302)

}

func (c *IndexController) DeleteType() {
	Id, err := c.GetInt("Id")
	if err != nil {
		beego.Info(err)
	}
	ArticleType := models.ArticleType{Id: Id}
	o := orm.NewOrm()
	err = o.Read(&ArticleType)
	if err != nil {
		beego.Info(err)
	} else {
		_, err = o.Delete(&ArticleType)
		if err != nil {
			beego.Info(err)
		} else {
			c.Redirect("/AddType", 302)
		}
	}
}

//logout control
func (c IndexController) Logout() {
	c.DelSession("userName")
	c.Redirect("/Login", 302)
}