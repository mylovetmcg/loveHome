package controllers

import (
	"github.com/astaxie/beego"
)

type HouseIndexController struct {
	beego.Controller
}

func (c *HouseIndexController)RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *HouseIndexController)GetHouseIndex(){
	resp := make(map[string]interface{})
	resp["errno"] = 0
	resp["errmsg"] = "查询成功"
	defer c.RetData(resp)
	return
}