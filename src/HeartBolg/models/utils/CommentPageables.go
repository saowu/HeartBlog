package utils

import (
	"HeartBolg/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type CommentPageables struct {
	//总页数
	PageCount int
	//当前页*M
	CurrentPage int
	//总条数
	TotalCount int
	//每页条数*M
	TotalNumber int
	//起始条数
	InitialNumber int
	//查询属性
	Fields string
	//数据
	Comments []models.Comment
}

//计算总页数
func (p *CommentPageables) SetPageCount() int {
	//初始化总页数
	if p.TotalCount != 0 {
		if (p.TotalCount % p.TotalNumber) == 0 {
			return p.TotalCount / p.TotalNumber
		} else {
			return (p.TotalCount / p.TotalNumber) + 1
		}
	}
	return 0
}

func (p *CommentPageables) SetPageableData() {
	i, err := orm.NewOrm().Raw("SELECT * FROM comment LIMIT ?,?", p.InitialNumber, p.TotalNumber).QueryRows(&p.Comments)
	if err == nil {
		fmt.Println("user nums: ", i)
	}
}

func (p *CommentPageables) SetInitialNumber() {
	p.InitialNumber = (p.CurrentPage - 1) * p.TotalNumber
}
