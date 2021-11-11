package service

import (
	"jrpc/resources/request"
	"jrpc/resources/response"
	"jrpc/resources/tools"
)

// News 消息服务
type News struct{}

// DingDing 钉钉消息服务
func (c *News) DingDing(params *request.ParamsDingDing, result *response.Result) error {
	dinging := &tools.DingDing{
		Params: request.ParamsDingDing{
			Host:      params.Host,
			Msg:       params.Msg,
			AtMobiles: params.AtMobiles,
		},
	}
	if params.Host.Address == "" {
		result.Error("参数格式不正确")
	}

	//发送
	res, err := dinging.Send()
	if err != nil {
		result.Error(err.Error())
	}

	result.Success("success", res)
	return nil
}
