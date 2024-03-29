package zentao_models

import "time"

// Leixing 类型表
type Leixing struct {
	Cloudname string  `json:"cloudname"` //类型
	Esti      float64 `json:"esti"`      //预计工时
	Cons      float64 `json:"cons"`      //实际工时
}

// Customer 客户表
type Customer struct {
	Customername string  `json:"customername"` //客户
	Esti         float64 `json:"esti"`         //预计工时
	Cons         float64 `json:"cons"`         //实际工时
}

// CustomerDetail 按类型取客户明细
type CustomerDetail struct {
	Customername string    `json:"customername"` //客户
	Id           int       `json:"id"`
	Titlename    string    `json:"titlename"` //问题标题
	Workdate     time.Time `json:"workdate"`  //处理时间
	Esti         float64   `json:"esti"`      //预计工时
	Cons         float64   `json:"cons"`      //实际工时
}

// LeixingDetail 按客户取类型明细
type LeixingDetail struct {
	Leixing   string    `json:"leixing"` //客户
	Id        int       `json:"id"`
	Titlename string    `json:"titlename"` //问题标题
	Workdate  time.Time `json:"workdate"`  //处理时间
	Esti      float64   `json:"esti"`      //预计工时
	Cons      float64   `json:"cons"`      //实际工时
}

type LeixingHeatmap struct {
	Cloudname string  `json:"cloudname"` //类型
	Esti      float64 `json:"esti"`      //预计工时
	Dateyear  int     `json:"dateyear"`  //年度
	Rk        int     `json:"rk"`        //类型排序
}

type CustomerHeatmap struct {
	Customername string  `json:"customername"` //客户
	Esti         float64 `json:"esti"`         //预计工时
	Dateyear     int     `json:"dateyear"`     //年度
	Rk           int     `json:"rk"`           //类型排序
}
