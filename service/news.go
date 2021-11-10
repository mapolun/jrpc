package service

import (
	request2 "jrpc/resources/request"
	response2 "jrpc/resources/response"
)

// News 消息服务
type News struct{}

// DingDing 钉钉消息服务
func (c *News) DingDing(params *request2.ParamsDingDing, result *response2.Result) error {
	result.Success("success", params)
	return nil
}
