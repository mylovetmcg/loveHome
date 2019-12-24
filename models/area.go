package models

type Area struct {
	Id     int      `json:"aid" description:"区域编号"`   //区域编号
	Name   string   `orm:"size(32)" json:"aname"`       //区域名字
	Houses []*House `orm:"reverse(many)" json:"houses"` //区域所有的房屋
}
