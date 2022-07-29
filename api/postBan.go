package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func PostBan(gid int64, uid int64, duiTime int64) string {

	// 结构体
	ban := struct {
		Group_id int64 `json:"group_id"`
		User_id  int64 `json:"user_id"`
		Duration int64 `json:"duration"`
	}{
		Group_id: gid,
		User_id:  uid,
		Duration: duiTime,
	}

	// 转换为json
	payload, _ := json.Marshal(&ban)

	postBan, _ := http.Post("http://localhost:5700/set_group_ban", "application/json", bytes.NewReader(payload))

	defer postBan.Body.Close()

	all, _ := ioutil.ReadAll(postBan.Body)

	fmt.Println("post返回信息：" + string(all))

	result := gjson.Get(string(all), "status").String()

	if result != "ok" {
		return "错误"
	}

	return "成员已经被禁言"
}
