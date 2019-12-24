package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){

	// set default database
	_ = orm.RegisterDataBase("default", beego.AppConfig.String("drivername"), "root:123456@tcp(127.0.0.1:3306)/loveHome?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User),new(OrderHouse),new(Facility),new(HouseImage),new(Area),new(House))

	// create table
	_ = orm.RunSyncdb("default", false, true)
}