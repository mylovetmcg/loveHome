package models

import "time"

const (
	ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"  //待接单
	ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT" //待支付
	ORDER_STATUS_PAID         = "PAID"         //已支付
	ORDER_STATUS_WAIT_COMMENT = "COMMENT"      //待评价
	ORDER_STATUS_COMPLETE     = "COMPLETE"     //已完成
	ORDER_STATUS_CANCELED     = "CONCELED"     //已取消
	ORDER_STATUS_REJECTED     = "REJECTED"     //已拒单
)

/* 订单 table_name = order */
type OrderHouse struct {
	Id          int       `json:"order_id"`               //订单编号
	User        *User     `orm:"rel(fk)" json:"user_id"`  //下单的用户编号
	House       *House    `orm:"rel(fk)" json:"house_id"` //预定的房间编号
	Begin_date  time.Time `orm:"type(datetime)"`          //预定的起始时间
	End_date    time.Time `orm:"type(datetime)"`          //预定的结束时间
	Days        int       `orm:"default(0)"`//预定总天数
	House_price int       `orm:"default(0)"`//房屋的单价
	Amount      int       `orm:"default(0)"`//订单总金额
	Status      string    `orm:"default(WAIT_ACCEPT)"` //订单状态
	Comment     string    `orm:"size(512)"`            //订单评论
	Ctime       time.Time `orm:"auto_now_add;type(datetime)" json:"ctime"`
}
