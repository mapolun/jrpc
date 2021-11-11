package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"jrpc/resources/request"
	untils2 "jrpc/resources/untils"
	"log"
	"net/url"
	"strconv"
	"time"
)

type DingDing struct {
	Params request.ParamsDingDing
}

type Err struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (d *DingDing) Send() (interface{}, error) {
	var data map[string]string
	var res map[string]interface{}
	var err error

	//组装发送数据包
	data, err = d.packData()
	if err != nil {
		return nil, err
	}

	//请求钉钉接口
	result := d.requestDingTalk(data)
	err = json.Unmarshal([]byte(result), &res)
	if err != nil {
		return nil, err
	}

	//返回结果
	return res, nil
}

func (d *DingDing) packData() (map[string]string, error) {
	var data map[string]string
	type text struct {
		Content string `json:"content"`
	}
	type at struct {
		AtMobiles []interface{} `json:"atMobiles"`
		IsAtAll   bool          `json:"isAtAll"`
	}

	textJson, err := json.Marshal(text{d.Params.Msg})
	if err != nil {
		return data, err
	}

	atJson, err := json.Marshal(at{d.Params.AtMobiles, false})
	if err != nil {
		return data, err
	}

	data = map[string]string{
		"msgtype": "text",
		"text":    string(textJson),
		"at":      string(atJson),
	}
	return data, nil
}

func (d *DingDing) requestDingTalk(data map[string]string) string {
	cli := untils2.NewHttpSend(d.getSendUrl())
	cli.SetSendType(untils2.SENDTYPE_JSON) //content-json
	cli.SetBody(data)
	resp, err := cli.Post()
	if err != nil {
		log.Println(err)
	}

	return string(resp)
}

func (d *DingDing) getSendUrl() string {
	var u string
	if d.Params.Host.Secret != "" {
		timeSpace := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
		sign := d.generateSign(timeSpace)
		u = d.Params.Host.Address + "&timestamp=" + timeSpace + "&sign=" + sign
	} else {
		u = d.Params.Host.Address
	}
	return u
}

func (d *DingDing) generateSign(timeSpace string) string {
	var sign string
	stringToSign := timeSpace + "\n" + d.Params.Host.Secret
	sign = ComputeHmacSha256(stringToSign, d.Params.Host.Secret)
	sign = base64.StdEncoding.EncodeToString([]byte(sign))
	sign = url.QueryEscape(sign)
	return sign
}

func ComputeHmacSha256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return string(h.Sum(nil))
}
