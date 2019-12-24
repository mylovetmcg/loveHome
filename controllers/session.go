package controllers

import (
	"github.com/astaxie/beego"
	"oys.org/loveHome/models"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController)RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *SessionController)GetSessionData(){
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	defer c.RetData(resp)

	name := c.GetSession("name")
	if name != nil {
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		user := models.User{
			Name: name.(string),
		}
		resp["data"] = user
	}
	return
}

func (c *SessionController)DeleteSessionData(){
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.RetData(resp)

	c.DelSession("name")
}