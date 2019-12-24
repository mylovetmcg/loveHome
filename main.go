package main

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	_ "oys.org/loveHome/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
	_ "oys.org/loveHome/models"
)

func main() {
	ignoreStaticPath()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

func ignoreStaticPath(){
	// 透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context){
	orpath := ctx.Request.URL.Path
	logs.Debug("request url: ", orpath)
	if strings.Index(orpath,"api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+orpath)
}