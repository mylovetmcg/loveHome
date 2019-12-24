package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"oys.org/loveHome/models"
)

type AreaController struct {
	beego.Controller
}

func (c *AreaController)RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AreaController)GetArea(){
	// 从session拿数据
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	// 从数据库中导入数据
	var areas []models.Area
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)

	if err != nil {
		resp["errno"] = 4001
		resp["errmsg"] = "查询失败"
		return
	}

	if num == 0 {
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}

	// 打包数据，打包成json，返回给前端
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = areas

	logs.Info(resp)
	return
}