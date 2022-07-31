package api

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

// GetSetu 涩图获取
func GetSetu() string {
	// API请求
	get, _ := http.Get("https://i.fs233.cc/setu?type=json")

	// 关闭流
	defer func() { _ = get.Body.Close() }()

	// 读取
	getRead, _ := ioutil.ReadAll(get.Body)

	msg := gjson.Get(string(getRead), "msg").String()

	return msg
}
