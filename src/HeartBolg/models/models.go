package models

import (
	"time"
)

//用户表
type User struct {
	Id       int32  `orm:"pk;auto"`
	UserName string `orm:"unique"`
	Password string
	Token    string
}

//文章类别
type Category struct {
	Id       int32      `json:"id" orm:"pk;auto"`
	Name     string     `json:"name" orm:"unique"`
	Articles []*Article `orm:"reverse(many)"`
}

//文章表
type Article struct {
	Id       int32      `json:"id"orm:"pk;auto"`
	Title    string     `json:"title"`
	Body     string     `json:"body"orm:"type(text)"`
	Time     time.Time  `json:"time"orm:"auto_now_add;type(datetime)"`
	Cover    string     `json:"cover"`
	Key      string     `json:"key"`
	Author   string     `json:"author"`
	Category *Category  `json:"category"orm:"rel(fk)"`
	Readers  *Reader    `json:"readers"orm:"rel(one)"`
	Comments []*Comment `json:"comments"orm:"reverse(many);on_delete(cascade)"`
}

//阅读表
type Reader struct {
	Id         int32    `json:"id" orm:"pk;auto"`
	Number     int      `json:"number" `
	Appreciate int      `json:"appreciate" `
	Articles   *Article `json:"articles" orm:"reverse(one)"`
}

//评论表
type Comment struct {
	Id       int32 `orm:"pk;auto"`
	Name     string
	Content  string `orm:"type(text)"`
	Email    string
	Time     time.Time `orm:"auto_now_add;type(datetime)"`
	Articles *Article  `orm:"rel(fk)"`
}

//文件表
type File struct {
	Id   int32 `orm:"pk;auto"`
	Name string
	Time time.Time `orm:"auto_now_add;type(datetime)"`
	Path string
}

//相册表
type Album struct {
	Id        int32 `orm:"pk;auto"`
	Title     string
	Content   string `orm:"type(text)"`
	PhotoPath string
	Time      time.Time `orm:"auto_now_add;type(datetime)"`
}

//推荐链接表
type Recommend struct {
	Id     int32  `json:"id" orm:"pk;auto"`
	Title  string `json:"title"`
	Link   string `json:"link"`
	Clicks int    `json:"clicks"`
}

//标签云
type TagCloud struct {
	Id    int32 `orm:"pk;auto"`
	Lable string
}

//关于
type About struct {
	Id       int32 `orm:"pk;auto"`
	Title    string
	Body     string    `orm:"type(text)"`
	Time     time.Time `orm:"auto_now_add;type(datetime)"`
	Photo    string
	Music    string
	BlogName string
}

//友情链接
type FriendshipLink struct {
	Id    int32 `orm:"pk;auto"`
	Title string
	Link  string
}
