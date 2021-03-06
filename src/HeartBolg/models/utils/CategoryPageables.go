package utils

import (
	"HeartBolg/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type CategoryPageables struct {
	//总页数
	PageCount int `json:"pageCount"`
	//当前页*M
	CurrentPage int `json:"currentPage"`
	//总条数
	TotalCount int `json:"totalCount"`
	//每页条数*M
	TotalNumber int `json:"totalNumber"`
	//起始条数
	InitialNumber int `json:"initialNumber"`
	//查询属性
	Fields string `json:"fields"`
	//数据
	Categorys []models.Category `json:"categorys"`
}

//计算总页数
func (p *CategoryPageables) SetPageCount() int {
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

func (p *CategoryPageables) SetPageableData() {
	i, err := orm.NewOrm().Raw("SELECT * FROM category LIMIT ?,?", p.InitialNumber, p.TotalNumber).QueryRows(&p.Categorys)
	if err == nil {
		fmt.Println("user nums: ", i)
	}
}

func (p *CategoryPageables) SetInitialNumber() {
	p.InitialNumber = (p.CurrentPage - 1) * p.TotalNumber
}
