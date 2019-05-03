package controllers

import (
	"HeartBolg/models"
	"HeartBolg/models/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const pagenum = 6

type DashboardController struct {
	beego.Controller
}

//控制台
func (d *DashboardController) Get() {
	o := orm.NewOrm()
	i, e := o.QueryTable("article").Count()
	if e != nil {
		fmt.Println(e)
	}
	i1, e1 := o.QueryTable("comment").Count()
	if e1 != nil {
		fmt.Println(e1)
	}
	i2, e2 := o.QueryTable("file").Count()
	if e2 != nil {
		fmt.Println(e2)
	}
	i3, e3 := o.QueryTable("friendship_link").Count()
	if e3 != nil {
		fmt.Println(e3)
	}
	d.Data["anum"] = i
	d.Data["cnum"] = i1
	d.Data["fnum"] = i2
	d.Data["flnum"] = i3
	d.TplName = "dashboard.html"
}

type ArticleController struct {
	beego.Controller
}

//文章管理
func (a *ArticleController) Get() {
	o := orm.NewOrm()
	//获取总数
	var article []models.Article
	var pageables utils.ArticlePageables
	o.QueryTable("article").All(&article)
	//分页
	pageables.TotalNumber = pagenum
	pageables.CurrentPage = 1
	pageables.SetInitialNumber()
	pageables.TotalCount = len(article)
	pageables.SetPageCount()
	pageables.SetPageableData()

	a.Data["pageables"] = pageables
	a.TplName = "article.html"
}

type CategoryController struct {
	beego.Controller
}

//类型管理
func (c *CategoryController) Get() {
	o := orm.NewOrm()
	//获取总数
	var category []models.Category
	var pageables utils.CategoryPageables
	o.QueryTable("category").All(&category)
	//分页
	pageables.TotalNumber = pagenum
	pageables.CurrentPage = 1
	pageables.SetInitialNumber()
	pageables.TotalCount = len(category)
	pageables.SetPageCount()
	pageables.SetPageableData()

	c.Data["pageables"] = pageables
	c.TplName = "articletype.html"
}

type LinkController struct {
	beego.Controller
}

//链接管理
func (l *LinkController) Get() {

	o := orm.NewOrm()
	var recommend []models.Recommend
	o.QueryTable("recommend").All(&recommend)

	//分页
	var pageables utils.RecommmendPageables
	pageables.TotalNumber = pagenum
	pageables.CurrentPage = 1
	pageables.SetInitialNumber()
	pageables.TotalCount = len(recommend)
	pageables.SetPageCount()
	pageables.SetPageableData()
	//========================================================
	var flink []models.FriendshipLink
	i, e := o.QueryTable("friendship_link").All(&flink)
	if e != nil {
		fmt.Println(i, e)
	}
	//分页
	var pageables1 utils.FriendshipLinkPageables
	pageables1.TotalNumber = pagenum
	pageables1.CurrentPage = 1
	pageables1.SetInitialNumber()
	pageables1.TotalCount = len(flink)
	pageables1.SetPageCount()
	pageables1.SetPageableData()

	l.Data["pageables"] = pageables
	l.Data["fpageables"] = pageables1
	fmt.Println(flink)
	l.TplName = "links.html"
}

type AlbumController struct {
	beego.Controller
}

//个人相册
func (a *AlbumController) Get() {
	o := orm.NewOrm()
	//获取总数
	var albums []models.Album
	var pageables utils.AlbumPageables
	o.QueryTable("album").All(&albums)
	//分页
	pageables.TotalNumber = pagenum
	pageables.CurrentPage = 1
	pageables.SetInitialNumber()
	pageables.TotalCount = len(albums)
	pageables.Fields = "album"
	pageables.SetPageCount()
	pageables.SetPageableData()
	//分页数据完整
	a.Data["pageables"] = pageables
	a.TplName = "album.html"
}

type FileController struct {
	beego.Controller
}

//文件管理
func (f *FileController) Get() {
	o := orm.NewOrm()
	//获取总数
	var files []models.File
	var pageables utils.FilePageables
	o.QueryTable("file").All(&files)
	//分页
	pageables.TotalNumber = pagenum
	pageables.CurrentPage = 1
	pageables.SetInitialNumber()
	pageables.TotalCount = len(files)
	pageables.SetPageCount()
	pageables.SetPageableData()
	//分页数据完整
	f.Data["pageables"] = pageables

	f.TplName = "files.html"
}

type SetMeController struct {
	beego.Controller
}

//关于我
func (s *SetMeController) Get() {
	s.TplName = "setme.html"
}

type ManagerTabCloudController struct {
	beego.Controller
}

//标签云
func (m *ManagerTabCloudController) Get() {
	o := orm.NewOrm()
	//获取总数
	var tagcloud []models.TagCloud
	var pageables utils.TagCloudPageables
	o.QueryTable("tag_cloud").All(&tagcloud)
	//分页
	pageables.TotalNumber = pagenum
	pageables.CurrentPage = 1
	pageables.SetInitialNumber()
	pageables.TotalCount = len(tagcloud)
	pageables.SetPageCount()
	pageables.SetPageableData()
	//分页数据完整
	m.Data["pageables"] = pageables
	m.TplName = "tagcloud.html"
}

type ManagerCommentController struct {
	beego.Controller
}

//评论管理
func (m *ManagerCommentController) Get() {
	o := orm.NewOrm()
	//获取总数
	var comment []models.Comment
	var pageables utils.CommentPageables
	o.QueryTable("comment").All(&comment)
	//分页
	pageables.TotalNumber = pagenum
	pageables.CurrentPage = 1
	pageables.SetInitialNumber()
	pageables.TotalCount = len(comment)
	pageables.Fields = "album"
	pageables.SetPageCount()
	pageables.SetPageableData()
	//分页数据完整
	m.Data["pageables"] = pageables

	m.TplName = "comment.html"
}
