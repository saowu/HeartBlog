package main

import (
	"HeartBolg/models"
	_ "HeartBolg/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(models.User), new(models.About), new(models.Category), new(models.Article), new(models.Reader), new(models.Comment), new(models.File), new(models.Album), new(models.Recommend), new(models.TagCloud), new(models.FriendshipLink))
	//mysql
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@tcp("+beego.AppConfig.String("mysqlurls")+":3306)/"+beego.AppConfig.String("mysqldb")+"?charset=utf8&loc=Local")
	//自动建表
	orm.RunSyncdb("default", false, true)
	//启动日志
	orm.Debug = false
}
func main() {
	beego.Run()
}
