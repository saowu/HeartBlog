package routers

import (
	"HeartBolg/controllers"
	"HeartBolg/models"
	"HeartBolg/models/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	// /admin and /index is static resource path
	beego.Get("/admin", func(context *context.Context) {
		context.Redirect(302, "/static/login.html")
	})

	beego.Get("/index", func(context *context.Context) {
		context.Redirect(302, "/static/index.html")
	})

	//this is Front end part path

	beego.Router("/", &controllers.IndexController{})

	beego.Router("/share", &controllers.ShareController{})

	beego.Router("/about", &controllers.AboutController{})

	beego.Router("/info", &controllers.InfoController{})

	beego.Router("/tagcloud", &controllers.TagCloudController{})

	beego.Router("/articletype", &controllers.ArticleTypeController{})

	beego.Router("/search", &controllers.SearchController{})

	beego.Router("/comment", &controllers.CommentController{})

	beego.Put("/praise", func(i *context.Context) {
		var reader models.Reader
		json.Unmarshal(i.Input.RequestBody, &reader)

		o := orm.NewOrm()
		if o.Read(&reader) == nil {
			reader.Appreciate = reader.Appreciate + 1
			if num, err := o.Update(&reader); err == nil {
				fmt.Println(num)
			}
		}
		bytes, e := json.Marshal(reader)
		if e != nil {
			log.Fatal(e)
			return
		}
		fmt.Println(string(bytes))
		i.Output.Body(bytes)
	})

	beego.Put("/recommendclick", func(i *context.Context) {
		var recommend models.Recommend
		json.Unmarshal(i.Input.RequestBody, &recommend)

		o := orm.NewOrm()
		if o.Read(&recommend) == nil {
			recommend.Clicks = recommend.Clicks + 1
			if num, err := o.Update(&recommend); err == nil {
				fmt.Println(num)
			}
		}
		bytes, e := json.Marshal(recommend)
		if e != nil {
			log.Fatal(e)
			return
		}
		i.Output.Body(bytes)
	})

	//=====================后台页面初始化===========================================

	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/link", &controllers.LinkController{})
	beego.Router("/album", &controllers.AlbumController{})
	beego.Router("/file", &controllers.FileController{})
	beego.Router("/setme", &controllers.SetMeController{})
	beego.Router("/tagcloud_manager", &controllers.ManagerTabCloudController{})
	beego.Router("/comments", &controllers.ManagerCommentController{})
	//========================类别管理========================================

	beego.Post("/categorys", func(i *context.Context) {
		var categorys models.Category
		json.Unmarshal(i.Input.RequestBody, &categorys)
		o := orm.NewOrm()
		insert, e := o.Insert(&categorys)
		if e != nil {
			fmt.Println(insert, e)
		} else {
			categorys.Id = int32(insert)
			bytes, _ := json.Marshal(categorys)
			i.Output.Body(bytes)
		}
	})
	beego.Delete("/categorys", func(i *context.Context) {
		var categorys models.Category
		json.Unmarshal(i.Input.RequestBody, &categorys)
		o := orm.NewOrm()
		_, e := o.Delete(&categorys)
		if e != nil {
			fmt.Println(e)
		} else {
			bytes, _ := json.Marshal(categorys)
			i.Output.Body(bytes)
		}
	})
	beego.Put("/categorys", func(i *context.Context) {
		var categorys models.Category
		json.Unmarshal(i.Input.RequestBody, &categorys)
		o := orm.NewOrm()
		insert, e := o.Update(&categorys)
		if e != nil {
			fmt.Println(insert, e)
		} else {
			categorys.Id = int32(insert)
			bytes, _ := json.Marshal(categorys)
			i.Output.Body(bytes)
		}
	})
	beego.Get("/categorys", func(i *context.Context) {
		var categorys models.Category
		id := i.Input.Query("id")
		atoi, _ := strconv.Atoi(id)
		categorys.Id = int32(atoi)
		o := orm.NewOrm()
		err := o.Read(&categorys)
		if err != nil {
			fmt.Println(err)
		} else {
			bytes, _ := json.Marshal(categorys)
			i.Output.Body(bytes)
		}
	})
	beego.Post("/cate_pageables", func(i *context.Context) {
		var pageables utils.CategoryPageables
		json.Unmarshal(i.Input.RequestBody, &pageables)
		o := orm.NewOrm()
		//获取总数
		count, _ := o.QueryTable("category").Count()
		//分页
		//pageables.TotalNumber = 6
		//pageables.CurrentPage = 1
		pageables.SetInitialNumber()
		pageables.TotalCount = int(count)
		pageables.SetPageCount()
		pageables.SetPageableData()
		bytes, _ := json.Marshal(pageables)
		i.Output.Body(bytes)
	})

	//========================文章管理===============================================

	beego.Post("/art_pageables", func(i *context.Context) {
		var pageables utils.ArticlePageables
		json.Unmarshal(i.Input.RequestBody, &pageables)
		o := orm.NewOrm()
		//获取总数
		count, _ := o.QueryTable("article").Count()
		//分页
		//pageables.TotalNumber = 6
		//pageables.CurrentPage = 1
		pageables.SetInitialNumber()
		pageables.TotalCount = int(count)
		pageables.SetPageCount()
		pageables.SetPageableData()
		bytes, _ := json.Marshal(pageables)
		i.Output.Body(bytes)

	})
	beego.Delete("/articles", func(i *context.Context) {
		var article models.Article
		json.Unmarshal(i.Input.RequestBody, &article)
		o := orm.NewOrm()
		_, e := o.Delete(&article)
		if e != nil {
			fmt.Println(e)
		} else {
			bytes, _ := json.Marshal(article)
			i.Output.Body(bytes)
		}
	})
	//上传封面图
	beego.Post("/cover_file", func(i *context.Context) {
		file, header, err := i.Request.FormFile("cover")
		if err != nil {
			fmt.Println(header.Filename, ":", err)
		}
		path := beego.AppConfig.String("filepath") + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day())
		all := os.MkdirAll(path, 0777)
		if all != nil {
			fmt.Println(all)
		}
		filepath := path + "/" + strconv.FormatInt(time.Now().Unix(), 10) + header.Filename
		defer file.Close()
		f, e := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if e != nil {
			fmt.Println(header.Filename, ":", e)
		}
		defer f.Close()
		io.Copy(f, file)
		var article models.Article
		article.Cover = filepath[1:len(filepath)]
		bytes, _ := json.Marshal(article)
		i.Output.Body(bytes)
	})
	beego.Get("/categoryall", func(i *context.Context) {
		var categoryall []models.Category
		o := orm.NewOrm()
		all, err := o.QueryTable("category").All(&categoryall)
		if err != nil {
			fmt.Println(err, all)
		} else {
			bytes, _ := json.Marshal(categoryall)
			i.Output.Body(bytes)
		}
	})
	//
	beego.Post("/articles", func(i *context.Context) {
		var a models.Article
		json.Unmarshal(i.Input.RequestBody, &a)

		var r models.Reader
		o := orm.NewOrm()
		i2, i3 := o.Insert(&r)
		if i3 != nil {
			fmt.Println(i2, i3)
		}
		a.Readers = &r
		fmt.Println(a.Category, a, a.Readers)

		insert, e := o.Insert(&a)
		if e != nil {
			fmt.Println(insert, e)
		}
		bytes, _ := json.Marshal(a)
		i.Output.Body(bytes)

	})

}
