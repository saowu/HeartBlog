package controllers

import (
	"HeartBolg/models"
	"HeartBolg/models/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//首页模板渲染
type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {
	o := orm.NewOrm()
	var articles []models.Article
	var about models.About
	var tags []models.TagCloud
	var recommends []models.Recommend
	var fslinks []models.FriendshipLink
	o.QueryTable("article").All(&articles)
	o.QueryTable("about").All(&about)
	o.QueryTable("tag_cloud").All(&tags)
	o.QueryTable("recommend").All(&recommends)
	o.QueryTable("friendship_link").All(&fslinks)

	//获取类型文章数
	var categorys []models.Category
	var m = make(map[string]int64)
	o.QueryTable("category").All(&categorys)
	for _, value := range categorys {
		v, e := o.QueryTable("article").Filter("Category", value.Id).Count()
		if e != nil {
			break
		}
		m[value.Name] = v
	}

	i.Data["articles"] = articles
	i.Data["about"] = about
	i.Data["tags"] = tags
	i.Data["categorys"] = m
	i.Data["recommends"] = recommends
	i.Data["fslinks"] = fslinks
	i.TplName = "index.html"
}

//share页模板渲染
type ShareController struct {
	beego.Controller
}

func (s *ShareController) Get() {
	o := orm.NewOrm()
	//获取总数
	var albums []models.Album
	var pageables utils.AlbumPageables
	o.QueryTable("album").All(&albums)
	//分页
	pageables.TotalNumber = 8
	pageables.CurrentPage = 1
	pageables.SetInitialNumber()
	pageables.TotalCount = len(albums)
	pageables.Fields = "album"
	pageables.SetPageCount()
	pageables.SetPageableData()
	//分页数据完整
	s.Data["pageables"] = pageables
	s.TplName = "share.html"
}

//about页模板渲染
type AboutController struct {
	beego.Controller
}

func (a *AboutController) Get() {
	o := orm.NewOrm()
	var about models.About
	o.QueryTable("about").All(&about)
	a.Data["about"] = about
	a.TplName = "about.html"
}

//Info页模板渲染
type InfoController struct {
	beego.Controller
}

func (i *InfoController) Get() {
	id := i.GetString("id")
	o := orm.NewOrm()
	var article models.Article
	var about models.About
	var tags []models.TagCloud
	var recommends []models.Recommend
	var rclicks []models.Recommend
	o.QueryTable("article").Filter("Id", id).One(&article)
	o.QueryTable("about").All(&about)
	o.QueryTable("tag_cloud").All(&tags)
	o.QueryTable("recommend").All(&recommends)
	//点击排行
	num, err := o.Raw("SELECT * FROM recommend ORDER BY clicks DESC").QueryRows(&rclicks)
	if err == nil {
		fmt.Println("user nums: ", num)
	}
	//获取类型文章数
	var categorys []models.Category
	var m = make(map[string]int64)
	o.QueryTable("category").All(&categorys)
	for _, value := range categorys {
		v, e := o.QueryTable("article").Filter("Category", value.Id).Count()
		if e != nil {
			break
		}
		m[value.Name] = v
	}
	//文章阅读量
	var reader models.Reader
	o.QueryTable("reader").Filter("Id", article.Readers.Id).One(&reader)
	//文章评论
	var comment []models.Comment
	o.QueryTable("comment").Filter("Articles", article.Id).All(&comment)
	//上一条
	var particle models.Article
	e := o.Raw("select * from article where id<? order by id desc limit 1", id).QueryRow(&particle)
	if e != nil {
		fmt.Println(e)
	}
	//下一条
	var narticle models.Article
	e1 := o.Raw("select * from article where id>? order by id asc limit 1", id).QueryRow(&narticle)
	if e1 != nil {
		fmt.Println(e1)
	}

	i.Data["comment"] = comment
	i.Data["reader"] = reader
	i.Data["article"] = article
	i.Data["about"] = about
	i.Data["tags"] = tags
	i.Data["categorys"] = m
	i.Data["recommends"] = recommends
	i.Data["rclicks"] = rclicks
	i.Data["particle"] = particle
	i.Data["narticle"] = narticle

	i.TplName = "info.html"
}
