package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"oys.org/loveHome/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController)RetData(resp map[string]interface{}){
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UserController)Reg(){
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	logs.Info(resp["mobile"])
	logs.Info(resp["password"])
	logs.Info(resp["sms_code"])

	o :=orm.NewOrm()
	user := &models.User{
		Name:          resp["mobile"].(string),
		Password_hash: resp["password"].(string),
		Mobile:        resp["mobile"].(string),
	}

	id, err := o.Insert(user)
	if err != nil {
		resp["errno"] = 4002
		resp["errmsg"] = "注册失败"
		return
	}

	logs.Info("reg success, id = ", id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"

	// 设置session
	c.SetSession("name", user.Name)
}