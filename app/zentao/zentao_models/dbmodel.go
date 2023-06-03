package zentao_models

// Leixing 类型表
type Leixing struct {
	Cloudname string  `json:"cloudname"` //类型
	Esti      float64 `json:"esti"`      //预计工时
	Cons      float64 `json:"cons"`      //实际工时
}
