package controllers

import (
	"HeartBolg/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type TagCloudController struct {
	beego.Controller
}

//标签云
func (t *TagCloudController) Get() {
	tag := t.GetString("tag")
	o := orm.NewOrm()
	var article []models.Article
	sql := "SELECT * FROM article  WHERE  `key`  LIKE  '%" + tag + "%'"
	i, e := o.Raw(sql).QueryRows(&article)
	if e != nil {
		fmt.Println("TagCloudController:", i, e)
	}

	//==========================================
	var about models.About
	var tags []models.TagCloud
	var recommends []models.Recommend
	var fslinks []models.FriendshipLink
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
	//=====================================
	t.Data["articles"] = article
	t.Data["about"] = about
	t.Data["tags"] = tags
	t.Data["categorys"] = m
	t.Data["recommends"] = recommends
	t.Data["fslinks"] = fslinks
	t.TplName = "index.html"
}

type ArticleTypeController struct {
	beego.Controller
}

//类别筛选文章
func (a *ArticleTypeController) Get() {
	name := a.GetString("name")
	o := orm.NewOrm()
	var article []models.Article
	var category models.Category
	o.QueryTable("category").Filter("Name", name).All(&category)
	i, e := o.QueryTable("article").Filter("Category", category.Id).All(&article)
	if e != nil {
		fmt.Println("ArticleTypeController:", i, e)
	}
	//==========================================
	var about models.About
	var tags []models.TagCloud
	var recommends []models.Recommend
	var fslinks []models.FriendshipLink
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
	//=====================================
	a.Data["articles"] = article
	a.Data["about"] = about
	a.Data["tags"] = tags
	a.Data["categorys"] = m
	a.Data["recommends"] = recommends
	a.Data["fslinks"] = fslinks
	a.TplName = "index.html"

}

type SearchController struct {
	beego.Controller
}

func (s *SearchController) Post() {
	body := s.GetString("body")
	o := orm.NewOrm()
	var article []models.Article
	sql := "SELECT * FROM article  WHERE  body  LIKE  '%" + body + "%'"
	i, e := o.Raw(sql).QueryRows(&article)
	if e != nil {
		fmt.Println("SearchController:", i, e)
	}

	//==========================================
	var about models.About
	var tabs []models.TagCloud
	var recommends []models.Recommend
	var fslinks []models.FriendshipLink
	o.QueryTable("about").All(&about)
	o.QueryTable("tag_cloud").All(&tabs)
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
	//=====================================
	s.Data["articles"] = article
	s.Data["about"] = about
	s.Data["tabs"] = tabs
	s.Data["categorys"] = m
	s.Data["recommends"] = recommends
	s.Data["fslinks"] = fslinks
	s.TplName = "index.html"

}

type CommentController struct {
	beego.Controller
}

//添加评论
func (c *CommentController) Post() {
	name := c.GetString("name")
	email := c.GetString("email")
	content := c.GetString("content")
	aid, _ := c.GetInt32("aid")
	o := orm.NewOrm()
	if name != "" && content != "" {
		var comments models.Comment
		var articles models.Article
		articles.Id = aid
		comments.Articles = &articles
		comments.Content = content
		comments.Email = email
		comments.Name = name
		comments.Time = time.Now()
		i, e := o.Insert(&comments)
		if e != nil {
			fmt.Println("CommentController:", i, e)
		}
	}
	//=============================================
	id := aid
	var article models.Article
	var about models.About
	var tabs []models.TagCloud
	var recommends []models.Recommend
	var rclicks []models.Recommend
	o.QueryTable("article").Filter("Id", id).One(&article)
	o.QueryTable("about").All(&about)
	o.QueryTable("tag_cloud").All(&tabs)
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
	c.Data["comment"] = comment
	c.Data["reader"] = reader
	c.Data["article"] = article
	c.Data["about"] = about
	c.Data["tabs"] = tabs
	c.Data["categorys"] = m
	c.Data["recommends"] = recommends
	c.Data["rclicks"] = rclicks
	c.Data["particle"] = particle
	c.Data["narticle"] = narticle
	c.TplName = "info.html"
}
