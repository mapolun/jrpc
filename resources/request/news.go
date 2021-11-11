package request

type Host struct {
	Address string
	Secret  string
}

type ParamsDingDing struct {
	Host      Host   `json:"host"`
	Msg       string `json:"msg"`
	AtMobiles []interface{}
}
