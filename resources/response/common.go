package response

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var Ok = 0
var No = 1

func (r *Result) Success(msg string, data interface{}) {
	r.Code = Ok
	r.Msg = msg
	r.Data = data
}

func (r *Result) Error(msg string) {
	r.Code = No
	r.Msg = msg
}
